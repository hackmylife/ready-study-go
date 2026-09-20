package koan

import (
	"testing"
)

func TestArea(t *testing.T) {
	equal(t, (Rectangle{Width: 3, Height: 4}).Area(), 12)
	equal(t, (Rectangle{}).Area(), 0)
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
