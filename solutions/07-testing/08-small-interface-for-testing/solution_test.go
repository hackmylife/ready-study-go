//go:build ignore

package koan

import (
	"errors"
	"strings"
	"testing"
)

type errorReader struct{ err error }

func (r errorReader) Read([]byte) (int, error) { return 0, r.err }
func TestCopyFirst(t *testing.T) {
	b, err := CopyFirst(strings.NewReader("abcdef"), 3)
	noError(t, err)
	equal(t, string(b), "abc")
	b, err = CopyFirst(strings.NewReader("a"), 3)
	noError(t, err)
	equal(t, string(b), "a")
	cause := errors.New("broken")
	_, err = CopyFirst(errorReader{cause}, 3)
	if !errors.Is(err, cause) {
		t.Fatal("error lost")
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
