package koan

import (
	"net/http/httptest"
	"testing"
)

func TestLimit(t *testing.T) {
	for _, tt := range []struct {
		q    string
		want int
	}{{"", 20}, {"?limit=1", 1}, {"?limit=100", 100}} {
		n, err := Limit(httptest.NewRequest("GET", "/"+tt.q, nil))
		noError(t, err)
		equal(t, n, tt.want)
	}
	for _, q := range []string{"?limit=", "?limit=0", "?limit=101", "?limit=x", "?limit=1&limit=2"} {
		_, err := Limit(httptest.NewRequest("GET", "/"+q, nil))
		wantError(t, err)
	}
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
