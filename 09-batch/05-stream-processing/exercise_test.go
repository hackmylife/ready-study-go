package koan

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"
)

type brokenWriter struct{}

func (brokenWriter) Write([]byte) (int, error) { return 0, io.ErrClosedPipe }
func TestTransform(t *testing.T) {
	var b bytes.Buffer
	noError(t, Transform(strings.NewReader("go\n猫\n"), &b))
	equal(t, b.String(), "GO\n猫\n")
	if !errors.Is(Transform(strings.NewReader("x"), brokenWriter{}), io.ErrClosedPipe) {
		t.Fatal("write error lost")
	}
	wantError(t, Transform(strings.NewReader(strings.Repeat("x", 128*1024)), io.Discard))
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
