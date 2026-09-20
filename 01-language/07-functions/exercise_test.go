package koan

import (
	"testing"
)

func TestApplyTwice(t *testing.T) {
	equal(t, ApplyTwice(3, func(v int) int { return v + 2 }), 7)
	calls := 0
	equal(t, ApplyTwice(2, func(v int) int { calls++; return v * 3 }), 18)
	equal(t, calls, 2)
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
