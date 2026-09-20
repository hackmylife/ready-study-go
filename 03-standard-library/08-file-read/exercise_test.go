package koan

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestRead(t *testing.T) {
	p := filepath.Join(t.TempDir(), "note.txt")
	noError(t, os.WriteFile(p, []byte("Go\n猫"), 0600))
	got, err := Read(p)
	noError(t, err)
	equal(t, got, "Go\n猫")
	_, err = Read(p + ".missing")
	if !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("want not-exist, got %v", err)
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
