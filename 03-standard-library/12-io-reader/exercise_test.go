package koan

import (
	"io"
	"strings"
	"testing"
)

type lastChunk struct{ done bool }

func (r *lastChunk) Read(p []byte) (int, error) {
	if r.done {
		return 0, io.EOF
	}
	r.done = true
	return copy(p, "Go"), io.EOF
}
func TestReadAll(t *testing.T) {
	b, err := ReadAll(&lastChunk{})
	noError(t, err)
	equal(t, string(b), "Go")
	b, err = ReadAll(strings.NewReader("猫"))
	noError(t, err)
	equal(t, string(b), "猫")
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
