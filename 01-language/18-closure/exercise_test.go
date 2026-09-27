package koan

import (
	"testing"
)

func TestCounter(t *testing.T) {
	first := Counter()
	equal(t, first(), 1)
	equal(t, first(), 2)
	equal(t, first(), 3)
	second := Counter()
	equal(t, second(), 1)
	equal(t, first(), 4)
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
