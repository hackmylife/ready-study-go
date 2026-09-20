package koan

import (
	"strings"
	"testing"
)

func TestDecoder(t *testing.T) {
	d := NewDecoder(strings.NewReader("a\nb"))
	line, ok := d.Next()
	equal(t, line, "a")
	equal(t, ok, true)
	line, ok = d.Next()
	equal(t, line, "b")
	equal(t, ok, true)
	_, ok = d.Next()
	equal(t, ok, false)
	noError(t, d.Err())
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
