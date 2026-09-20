package koan

import (
	"testing"
)

func TestReverse(t *testing.T) {
	equal(t, Reverse("Go猫"), "猫oG")
	equal(t, Reverse("a🙂b"), "b🙂a")
	equal(t, Reverse(""), "")
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
