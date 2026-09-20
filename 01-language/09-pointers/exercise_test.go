package koan

import (
	"testing"
)

func TestIncrement(t *testing.T) {
	n := 4
	equal(t, Increment(&n), true)
	equal(t, n, 5)
	equal(t, Increment(nil), false)
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
