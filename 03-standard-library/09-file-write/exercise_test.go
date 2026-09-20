package koan

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWrite(t *testing.T) {
	p := filepath.Join(t.TempDir(), "out")
	noError(t, Write(p, "long content"))
	noError(t, Write(p, "Go"))
	b, err := os.ReadFile(p)
	noError(t, err)
	equal(t, string(b), "Go")
	info, err := os.Stat(p)
	noError(t, err)
	equal(t, info.Mode().Perm()&0077, os.FileMode(0))
	wantError(t, Write(filepath.Join(p, "missing"), "x"))
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
