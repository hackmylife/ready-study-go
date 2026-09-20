package koan

import (
	"net/http"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	before := http.DefaultClient.Timeout
	c, err := NewClient(2 * time.Second)
	noError(t, err)
	equal(t, c.Timeout, 2*time.Second)
	if c == http.DefaultClient {
		t.Fatal("do not mutate the shared default")
	}
	equal(t, http.DefaultClient.Timeout, before)
	for _, d := range []time.Duration{0, -1} {
		_, err := NewClient(d)
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
