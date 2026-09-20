package koan

import (
	"testing"
)

func TestCount(t *testing.T) { equal(t, Count(1000), 1000); equal(t, Count(0), 0) }
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
