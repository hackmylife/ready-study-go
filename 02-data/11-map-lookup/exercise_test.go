package koan

import (
	"testing"
)

func TestStock(t *testing.T) {
	equal(t, Stock(map[string]int{"tea": 4}, "tea"), 4)
	equal(t, Stock(nil, "tea"), 0)
	equal(t, Stock(map[string]int{}, "coffee"), 0)
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
