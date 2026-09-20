package koan

import (
	"net/http/httptest"
	"testing"
)

func TestRoutes(t *testing.T) {
	h := Routes()
	for _, tt := range []struct {
		method, path string
		status       int
	}{{"GET", "/health", 200}, {"POST", "/health", 405}, {"GET", "/missing", 404}} {
		w := httptest.NewRecorder()
		h.ServeHTTP(w, httptest.NewRequest(tt.method, tt.path, nil))
		equal(t, w.Code, tt.status)
		if tt.status == 200 {
			equal(t, w.Body.String(), "ok")
		}
	}
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
