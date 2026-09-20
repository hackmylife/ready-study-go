package koan

import (
	"encoding/json"
	"net/http/httptest"
	"reflect"
	"strings"
	"sync"
	"testing"
)

func TestNotes(t *testing.T) {
	h := NewHandler()
	w := httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest("POST", "/notes", strings.NewReader(`{"text":"learn Go"}`)))
	equal(t, w.Code, 201)
	equal(t, w.Header().Get("Content-Type"), "application/json")
	var created Note
	noError(t, json.Unmarshal(w.Body.Bytes(), &created))
	equal(t, created.Text, "learn Go")
	location := w.Header().Get("Location")
	if location == "" {
		t.Fatal("missing Location")
	}
	w = httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest("GET", location, nil))
	equal(t, w.Code, 200)
	var found Note
	noError(t, json.Unmarshal(w.Body.Bytes(), &found))
	same(t, found, created)
	w = httptest.NewRecorder()
	NewHandler().ServeHTTP(w, httptest.NewRequest("GET", location, nil))
	equal(t, w.Code, 404)
}
func TestNotesInvalid(t *testing.T) {
	h := NewHandler()
	for _, body := range []string{`{"text":" "}`, `{"other":1}`, `{"text":"x"} {}`, `null`, `{`} {
		w := httptest.NewRecorder()
		h.ServeHTTP(w, httptest.NewRequest("POST", "/notes", strings.NewReader(body)))
		equal(t, w.Code, 400)
	}
}
func TestNotesConcurrent(t *testing.T) {
	h := NewHandler()
	ids := make(chan int, 20)
	errs := make(chan string, 20)
	var wg sync.WaitGroup
	for range 20 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			w := httptest.NewRecorder()
			h.ServeHTTP(w, httptest.NewRequest("POST", "/notes", strings.NewReader(`{"text":"x"}`)))
			var n Note
			if w.Code != 201 || json.Unmarshal(w.Body.Bytes(), &n) != nil {
				errs <- w.Body.String()
				return
			}
			ids <- n.ID
		}()
	}
	wg.Wait()
	close(ids)
	close(errs)
	for err := range errs {
		t.Error(err)
	}
	seen := make(map[int]bool)
	for id := range ids {
		if seen[id] {
			t.Error("duplicate id")
		}
		seen[id] = true
	}
	equal(t, len(seen), 20)
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
