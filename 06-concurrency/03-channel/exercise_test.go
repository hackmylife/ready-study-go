package koan

import (
	"testing"
)

func TestSum(t *testing.T) {
	in := make(chan int, 3)
	in <- 2
	in <- 0
	in <- 5
	close(in)
	equal(t, Sum(in), 7)
	empty := make(chan int)
	close(empty)
	equal(t, Sum(empty), 0)
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
