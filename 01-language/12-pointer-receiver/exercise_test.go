package koan

import (
	"testing"
)

func TestCounter(t *testing.T) {
	var c Counter
	c.Add(3)
	c.Add(-1)
	equal(t, c.Value(), 2)
	other := c
	other.Add(4)
	equal(t, c.Value(), 2)
	equal(t, other.Value(), 6)
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
