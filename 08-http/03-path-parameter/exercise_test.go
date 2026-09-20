package koan

import (
	"net/http/httptest"
	"testing"
)

func TestPath(t *testing.T) {
	for _, id := range []string{"u1", "alice"} {
		w := httptest.NewRecorder()
		Routes().ServeHTTP(w, httptest.NewRequest("GET", "/users/"+id, nil))
		equal(t, w.Code, 200)
		equal(t, w.Body.String(), id)
	}
	w := httptest.NewRecorder()
	Routes().ServeHTTP(w, httptest.NewRequest("GET", "/users/a/extra", nil))
	equal(t, w.Code, 404)
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
