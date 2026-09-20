package koan

import (
	"bytes"
	"errors"
	"io"
	"testing"
)

type short struct{}

func (short) Write(p []byte) (int, error) { return 1, io.ErrShortWrite }
func TestCounting(t *testing.T) {
	var b bytes.Buffer
	w := &CountingWriter{Writer: &b}
	n, err := w.Write([]byte("Go"))
	noError(t, err)
	equal(t, n, 2)
	_, err = w.Write([]byte("!"))
	noError(t, err)
	equal(t, w.Bytes, 3)
	equal(t, b.String(), "Go!")
	w = &CountingWriter{Writer: short{}}
	n, err = w.Write([]byte("abc"))
	equal(t, n, 1)
	equal(t, w.Bytes, 1)
	if !errors.Is(err, io.ErrShortWrite) {
		t.Fatal(err)
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
