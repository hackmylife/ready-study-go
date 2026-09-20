package koan

import (
	"testing"
)

func TestByteLength(t *testing.T) {
	equal(t, ByteLength("Go"), 2)
	equal(t, ByteLength("猫"), 3)
	equal(t, ByteLength(""), 0)
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
