package koan

import (
	"net/http/httptest"
	"testing"
)

func TestCreated(t *testing.T) {
	w := httptest.NewRecorder()
	Created(w, httptest.NewRequest("POST", "/users", nil))
	equal(t, w.Code, 201)
	equal(t, w.Header().Get("Location"), "/users/u1")
	equal(t, w.Body.Len(), 0)
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
