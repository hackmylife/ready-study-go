package koan

import (
	"fmt"
	"testing"
)

func TestAge(t *testing.T) {
	for _, n := range []int{0, 42, 150} {
		noError(t, ValidateAge(n))
	}
	for _, n := range []int{-1, 151} {
		err := ValidateAge(n)
		wantError(t, err)
		equal(t, err.Error(), fmt.Sprintf("age out of range: %d", n))
	}
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
