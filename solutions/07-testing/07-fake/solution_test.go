//go:build ignore

package koan

import (
	"errors"
	"testing"
)

type fakeNotifier struct {
	to, body string
	err      error
}

func (f *fakeNotifier) Notify(to, body string) error {
	f.to = to
	f.body = body
	return f.err
}
func TestSend(t *testing.T) {
	f := &fakeNotifier{}
	noError(t, Send(f, "aki@example.test"))
	equal(t, f.to, "aki@example.test")
	equal(t, f.body, "Welcome")
	f.err = errors.New("offline")
	if !errors.Is(Send(f, "x"), f.err) {
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
