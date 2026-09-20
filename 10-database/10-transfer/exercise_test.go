package koan

import (
	"context"
	"database/sql"
	"errors"
	"ready-study-go/internal/dbtest"
	"testing"
	"time"
)

func TestTransfer(t *testing.T) {
	db := dbtest.Open(t)
	dbtest.Accounts(t, db)
	noError(t, Transfer(context.Background(), db, "a", "b", 30))
	equal(t, dbtest.Balance(t, db, "a"), int64(70))
	equal(t, dbtest.Balance(t, db, "b"), int64(80))
	for _, tt := range []struct {
		from, to string
		amount   int64
		cause    error
	}{{"a", "b", 71, ErrFunds}, {"a", "a", 1, ErrInvalid}, {"a", "b", 0, ErrInvalid}, {"a", "b", -1, ErrInvalid}, {"a", "missing", 1, sql.ErrNoRows}} {
		err := Transfer(context.Background(), db, tt.from, tt.to, tt.amount)
		if !errors.Is(err, tt.cause) {
			t.Fatalf("got %v, want %v", err, tt.cause)
		}
	}
	equal(t, dbtest.Balance(t, db, "a"), int64(70))
	equal(t, dbtest.Balance(t, db, "b"), int64(80))
}
func TestTransferConcurrent(t *testing.T) {
	db := dbtest.Open(t)
	dbtest.Accounts(t, db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	errs := make(chan error, 2)
	for range 2 {
		go func() { errs <- Transfer(ctx, db, "a", "b", 80) }()
	}
	success, insufficient := 0, 0
	for range 2 {
		err := <-errs
		if err == nil {
			success++
		} else if errors.Is(err, ErrFunds) {
			insufficient++
		} else {
			t.Fatal(err)
		}
	}
	equal(t, success, 1)
	equal(t, insufficient, 1)
	equal(t, dbtest.Balance(t, db, "a"), int64(20))
	equal(t, dbtest.Balance(t, db, "b"), int64(130))
}
func TestTransferRollbackAndOverflow(t *testing.T) {
	db := dbtest.Open(t)
	dbtest.Accounts(t, db)
	dbtest.Exec(t, db, `ALTER TABLE accounts ADD CHECK(id <> 'b' OR balance <= 60)`)
	wantError(t, Transfer(context.Background(), db, "a", "b", 20))
	equal(t, dbtest.Balance(t, db, "a"), int64(100))
	equal(t, dbtest.Balance(t, db, "b"), int64(50))
}

func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
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
