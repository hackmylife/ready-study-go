package koan

import (
	"context"
	"net/http/httptest"
	"testing"
)

func TestContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	calls := 0
	h := Handler(func(got context.Context) (string, error) { calls++; equal(t, got, ctx); return "", got.Err() })
	w := httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest("GET", "/", nil).WithContext(ctx))
	equal(t, w.Code, 503)
	equal(t, calls, 1)
	w = httptest.NewRecorder()
	Handler(func(context.Context) (string, error) { return "ok", nil }).ServeHTTP(w, httptest.NewRequest("GET", "/", nil))
	equal(t, w.Code, 200)
	equal(t, w.Body.String(), "ok")
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
