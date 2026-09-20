package koan

import (
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	w := httptest.NewRecorder()
	Hello(w, httptest.NewRequest("GET", "/", nil))
	equal(t, w.Code, 200)
	equal(t, w.Body.String(), "Hello, Go\n")
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
