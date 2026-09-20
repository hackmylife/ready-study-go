package koan

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreate(t *testing.T) {
	for _, tt := range []struct {
		body   string
		status int
	}{{`{"name":"Aki"}`, 204}, {`{"name":""}`, 400}, {`null`, 400}, {`{"name":"x","extra":1}`, 400}, {`{"name":"x"} {}`, 400}, {`{"name":"` + strings.Repeat("x", 1024) + `"}`, 400}, {`{`, 400}} {
		w := httptest.NewRecorder()
		Create(w, httptest.NewRequest("POST", "/", strings.NewReader(tt.body)))
		equal(t, w.Code, tt.status)
	}
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
