package koan

import (
	"testing"
)

func TestCountPositive(t *testing.T) {
	equal(t, CountPositive([]int{-1, 0, 3, 4}), 2)
	equal(t, CountPositive(nil), 0)
	equal(t, CountPositive([]int{-3}), 0)
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
