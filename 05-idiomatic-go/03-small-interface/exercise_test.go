package koan

import (
	"bytes"
	"io"
	"testing"
)

func TestSend(t *testing.T) {
	send, ok := any(Send).(func(io.Writer, string) error)
	if !ok {
		t.Fatal("Send must accept io.Writer")
	}
	var b bytes.Buffer
	noError(t, send(&b, "Go"))
	equal(t, b.String(), "Go")
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
