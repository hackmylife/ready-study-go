package koan

import (
	"testing"
)

func TestNormalize(t *testing.T) {
	equal(t, Normalize("  Go\t is\n fun  "), "Go is fun")
	equal(t, Normalize("\u3000猫\u3000犬"), "猫 犬")
	equal(t, Normalize(" "), "")
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
