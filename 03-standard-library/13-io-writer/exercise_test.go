package koan

import (
	"bytes"
	"errors"
	"testing"
)

var errWrite = errors.New("write failed")

type failingWriter struct{}

func (failingWriter) Write([]byte) (int, error) { return 0, errWrite }
func TestGreeting(t *testing.T) {
	var b bytes.Buffer
	noError(t, WriteGreeting(&b, "猫"))
	equal(t, b.String(), "Hello, 猫\n")
	if !errors.Is(WriteGreeting(failingWriter{}, "A"), errWrite) {
		t.Fatal("write error lost")
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
