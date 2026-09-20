package koan

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddleware(t *testing.T) {
	for _, tt := range []struct{ in, want string }{{"r-1", "r-1"}, {"", "unknown"}} {
		calls := 0
		h := RequestID(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { calls++; w.WriteHeader(204) }))
		r := httptest.NewRequest("GET", "/", nil)
		r.Header.Set("X-Request-ID", tt.in)
		w := httptest.NewRecorder()
		h.ServeHTTP(w, r)
		equal(t, w.Code, 204)
		equal(t, w.Header().Get("X-Request-ID"), tt.want)
		equal(t, calls, 1)
	}
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
