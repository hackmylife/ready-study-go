package koan

import (
	"testing"
	"time"
)

func TestOptions(t *testing.T) {
	c, err := NewClient()
	noError(t, err)
	equal(t, c.Timeout, 5*time.Second)
	c, err = NewClient(WithTimeout(time.Second), WithTimeout(2*time.Second))
	noError(t, err)
	equal(t, c.Timeout, 2*time.Second)
	for _, d := range []time.Duration{0, -1} {
		_, err := NewClient(WithTimeout(d))
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
