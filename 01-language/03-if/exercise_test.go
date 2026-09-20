package koan

import (
	"testing"
)

func TestClamp(t *testing.T) {
	for _, tt := range []struct{ in, want int }{{-1, 0}, {0, 0}, {4, 4}, {10, 10}, {11, 10}} {
		equal(t, Clamp(tt.in, 0, 10), tt.want)
	}
	equal(t, Clamp(9, 3, 3), 3)
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
