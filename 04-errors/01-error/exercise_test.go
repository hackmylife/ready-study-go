package koan

import (
	"testing"
)

func TestSqrt(t *testing.T) {
	for _, tt := range []struct{ n, want float64 }{{0, 0}, {9, 3}, {0.25, 0.5}} {
		got, err := Sqrt(tt.n)
		noError(t, err)
		equal(t, got, tt.want)
	}
	_, err := Sqrt(-1)
	wantError(t, err)
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
func noError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
func wantError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("expected an error")
	}
}
