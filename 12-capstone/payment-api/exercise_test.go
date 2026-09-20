package koan

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"math"
	"net/http"
	"net/http/httptest"
	"ready-study-go/internal/dbtest"
	"reflect"
	"strings"
	"testing"
	"time"
)

func setup(t *testing.T) (*Store, http.Handler) {
	t.Helper()
	db := dbtest.Open(t)
	noError(t, Migrate(context.Background(), db))
	store := &Store{DB: db}
	return store, Routes(store, slog.New(slog.NewJSONHandler(io.Discard, nil)))
}
func request(t *testing.T, h http.Handler, method, path, body, key string) *httptest.ResponseRecorder {
	t.Helper()
	r := httptest.NewRequest(method, path, strings.NewReader(body))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Idempotency-Key", key)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	return w
}
func account(t *testing.T, s *Store, name string, balance int64) Account {
	t.Helper()
	a, err := s.CreateAccount(context.Background(), name, balance)
	noError(t, err)
	if a.ID <= 0 {
		t.Fatal("account ID must be positive")
	}
	return a
}

func TestAccountAPI(t *testing.T) {
	_, h := setup(t)
	w := request(t, h, "POST", "/accounts", `{"name":" Aki ","balance":100}`, "")
	equal(t, w.Code, 201)
	equal(t, w.Header().Get("Content-Type"), "application/json")
	var a Account
	noError(t, json.Unmarshal(w.Body.Bytes(), &a))
	equal(t, a.Name, "Aki")
	equal(t, a.Balance, int64(100))
	location := w.Header().Get("Location")
	if location == "" {
		t.Fatal("missing Location")
	}
	w = request(t, h, "GET", location, "", "")
	equal(t, w.Code, 200)
	var got Account
	noError(t, json.Unmarshal(w.Body.Bytes(), &got))
	same(t, got, a)
}

func TestAccountValidation(t *testing.T) {
	_, h := setup(t)
	for _, body := range []string{`{"name":" ","balance":0}`, `{"name":"x","balance":-1}`, `{"name":"x","balance":0.5}`, `{"name":"x","extra":1}`, `{"name":"x"} {}`, `null`, `{`, `{"name":"` + strings.Repeat("x", 4096) + `"}`} {
		w := request(t, h, "POST", "/accounts", body, "")
		equal(t, w.Code, 400)
	}
	for _, tt := range []struct {
		path   string
		status int
	}{{"/accounts/no", 400}, {"/accounts/0", 400}, {"/accounts/999999", 404}, {"/transfers/no", 400}, {"/transfers/999999", 404}} {
		w := request(t, h, "GET", tt.path, "", "")
		equal(t, w.Code, tt.status)
	}
}

func TestTransferAPIAndReplay(t *testing.T) {
	s, h := setup(t)
	a, b := account(t, s, "A", 100), account(t, s, "B", 50)
	body := fmt.Sprintf(`{"from_id":%d,"to_id":%d,"amount":30}`, a.ID, b.ID)
	w := request(t, h, "POST", "/transfers", body, "order-1")
	equal(t, w.Code, 201)
	var tr Transfer
	noError(t, json.Unmarshal(w.Body.Bytes(), &tr))
	equal(t, tr.FromID, a.ID)
	equal(t, tr.ToID, b.ID)
	equal(t, tr.Amount, int64(30))
	location := w.Header().Get("Location")
	if location == "" {
		t.Fatal("missing Location")
	}
	w = request(t, h, "POST", "/transfers", body, "order-1")
	equal(t, w.Code, 200)
	var replay Transfer
	noError(t, json.Unmarshal(w.Body.Bytes(), &replay))
	same(t, replay, tr)
	w = request(t, h, "GET", location, "", "")
	equal(t, w.Code, 200)
	var found Transfer
	noError(t, json.Unmarshal(w.Body.Bytes(), &found))
	same(t, found, tr)
	got, err := s.Account(context.Background(), a.ID)
	noError(t, err)
	equal(t, got.Balance, int64(70))
	got, err = s.Account(context.Background(), b.ID)
	noError(t, err)
	equal(t, got.Balance, int64(80))
	changed := fmt.Sprintf(`{"from_id":%d,"to_id":%d,"amount":31}`, a.ID, b.ID)
	w = request(t, h, "POST", "/transfers", changed, "order-1")
	equal(t, w.Code, 409)
	var count int
	noError(t, s.DB.QueryRow(`SELECT count(*) FROM transfers`).Scan(&count))
	equal(t, count, 1)
}

func TestTransferValidationAndFunds(t *testing.T) {
	s, h := setup(t)
	a, b := account(t, s, "A", 100), account(t, s, "B", 50)
	for _, tt := range []struct {
		from, to, amount int64
		key              string
		status           int
	}{{a.ID, b.ID, 101, "funds", 409}, {a.ID, a.ID, 1, "same", 400}, {a.ID, b.ID, 0, "zero", 400}, {a.ID, b.ID, -1, "negative", 400}, {a.ID, 999999, 1, "missing", 404}, {a.ID, b.ID, 1, "", 400}} {
		body := fmt.Sprintf(`{"from_id":%d,"to_id":%d,"amount":%d}`, tt.from, tt.to, tt.amount)
		w := request(t, h, "POST", "/transfers", body, tt.key)
		equal(t, w.Code, tt.status)
	}
	got, err := s.Account(context.Background(), a.ID)
	noError(t, err)
	equal(t, got.Balance, int64(100))
	var count int
	noError(t, s.DB.QueryRow(`SELECT count(*) FROM transfers`).Scan(&count))
	equal(t, count, 0)
}

func TestConcurrentOverspend(t *testing.T) {
	s, _ := setup(t)
	a, b := account(t, s, "A", 100), account(t, s, "B", 0)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	errs := make(chan error, 2)
	for i := range 2 {
		go func() { _, _, err := s.Send(ctx, fmt.Sprintf("key-%d", i), a.ID, b.ID, 80); errs <- err }()
	}
	success, failed := 0, 0
	for range 2 {
		err := <-errs
		if err == nil {
			success++
		} else if errors.Is(err, ErrFunds) {
			failed++
		} else {
			t.Fatal(err)
		}
	}
	equal(t, success, 1)
	equal(t, failed, 1)
	got, err := s.Account(ctx, a.ID)
	noError(t, err)
	equal(t, got.Balance, int64(20))
	got, err = s.Account(ctx, b.ID)
	noError(t, err)
	equal(t, got.Balance, int64(80))
}

func TestConcurrentSameKey(t *testing.T) {
	s, _ := setup(t)
	a, b := account(t, s, "A", 100), account(t, s, "B", 0)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	type outcome struct {
		tr      Transfer
		created bool
		err     error
	}
	results := make(chan outcome, 2)
	for range 2 {
		go func() { tr, created, err := s.Send(ctx, "retry", a.ID, b.ID, 80); results <- outcome{tr, created, err} }()
	}
	first, second := <-results, <-results
	noError(t, first.err)
	noError(t, second.err)
	same(t, first.tr, second.tr)
	if first.created == second.created {
		t.Fatal("exactly one request must create the transfer")
	}
	got, err := s.Account(ctx, a.ID)
	noError(t, err)
	equal(t, got.Balance, int64(20))
}

func TestTransferRollback(t *testing.T) {
	s, _ := setup(t)
	a, b := account(t, s, "A", 100), account(t, s, "B", 0)
	dbtest.Exec(t, s.DB, `ALTER TABLE transfers ADD CHECK(amount<=10)`)
	_, _, err := s.Send(context.Background(), "fail", a.ID, b.ID, 20)
	wantError(t, err)
	got, err := s.Account(context.Background(), a.ID)
	noError(t, err)
	equal(t, got.Balance, int64(100))
	got, err = s.Account(context.Background(), b.ID)
	noError(t, err)
	equal(t, got.Balance, int64(0))
	var count int
	noError(t, s.DB.QueryRow(`SELECT count(*) FROM transfers`).Scan(&count))
	equal(t, count, 0)
}

func TestCancellationAndOverflow(t *testing.T) {
	s, _ := setup(t)
	a, b := account(t, s, "A", 100), account(t, s, "B", math.MaxInt64)
	_, _, err := s.Send(context.Background(), "overflow", a.ID, b.ID, 1)
	if !errors.Is(err, ErrInvalid) {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, _, err = s.Send(ctx, "canceled", a.ID, b.ID, 1)
	if !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}
	_, err = s.Account(ctx, a.ID)
	if !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}
}

func TestInternalErrorResponse(t *testing.T) {
	s, _ := setup(t)
	var logs bytes.Buffer
	h := Routes(s, slog.New(slog.NewJSONHandler(&logs, nil)))
	dbtest.Exec(t, s.DB, `DROP TABLE transfers`)
	w := request(t, h, "GET", "/transfers/1", "", "")
	equal(t, w.Code, 500)
	var body map[string]string
	noError(t, json.Unmarshal(w.Body.Bytes(), &body))
	same(t, body, map[string]string{"error": "internal_error"})
	if !strings.Contains(logs.String(), "request failed") {
		t.Fatal("internal failure was not logged")
	}
}

func TestConflictingKeysAcrossDifferentAccounts(t *testing.T) {
	s, _ := setup(t)
	a, b := account(t, s, "A", 100), account(t, s, "B", 0)
	c, d := account(t, s, "C", 100), account(t, s, "D", 0)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	errs := make(chan error, 2)
	go func() { _, _, err := s.Send(ctx, "shared-key", a.ID, b.ID, 10); errs <- err }()
	go func() { _, _, err := s.Send(ctx, "shared-key", c.ID, d.ID, 20); errs <- err }()
	success, conflict := 0, 0
	for range 2 {
		err := <-errs
		if err == nil {
			success++
		} else if errors.Is(err, ErrConflict) {
			conflict++
		} else {
			t.Fatalf("want success or idempotency conflict, got %v", err)
		}
	}
	equal(t, success, 1)
	equal(t, conflict, 1)
	var tr Transfer
	noError(t, s.DB.QueryRow(`SELECT id,from_id,to_id,amount FROM transfers`).Scan(&tr.ID, &tr.FromID, &tr.ToID, &tr.Amount))
	for _, initial := range []Account{a, b, c, d} {
		want := initial.Balance
		if initial.ID == tr.FromID {
			want -= tr.Amount
		}
		if initial.ID == tr.ToID {
			want += tr.Amount
		}
		got, err := s.Account(ctx, initial.ID)
		noError(t, err)
		equal(t, got.Balance, want)
	}
}

func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
func same(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v; want %#v", got, want)
	}
}
func noError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
func wantError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("expected an error")
	}
}
