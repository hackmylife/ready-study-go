package koan

import (
	"testing"
)

func TestSwap(t *testing.T) {
	for _, tt := range [][2]int{{2, 9}, {-3, 0}, {4, 4}} {
		a, b := Swap(tt[0], tt[1])
		equal(t, a, tt[1])
		equal(t, b, tt[0])
	}
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
