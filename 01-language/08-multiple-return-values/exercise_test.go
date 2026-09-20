package koan

import (
	"testing"
)

func TestDivide(t *testing.T) {
	for _, tt := range []struct {
		a, b, q int
		ok      bool
	}{{7, 2, 3, true}, {-6, 2, -3, true}, {0, 2, 0, true}, {7, 0, 0, false}} {
		q, ok := Divide(tt.a, tt.b)
		equal(t, q, tt.q)
		equal(t, ok, tt.ok)
	}
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
