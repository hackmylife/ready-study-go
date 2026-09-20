package koan

import (
	"testing"
)

func TestSumTo(t *testing.T) {
	for _, tt := range []struct{ n, want int }{{-2, 0}, {0, 0}, {1, 1}, {4, 10}, {8, 36}} {
		equal(t, SumTo(tt.n), tt.want)
	}
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
