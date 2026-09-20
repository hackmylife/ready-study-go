//go:build ignore

package koan

import (
	"testing"
)

func TestAbs(t *testing.T) {
	equal(t, Abs(5), 5)
	equal(t, Abs(-5), 5)
	equal(t, Abs(0), 0)
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
