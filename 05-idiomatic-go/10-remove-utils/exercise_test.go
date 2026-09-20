package koan

import (
	"errors"
	"os"
	"testing"
)

func TestCode(t *testing.T) {
	equal(t, NormalizeCode("  go-42 "), "GO-42")
	equal(t, NormalizeCode(""), "")
	if _, err := os.Stat("utils"); !errors.Is(err, os.ErrNotExist) {
		t.Fatal("move the implementation and remove the utils package")
	}
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
