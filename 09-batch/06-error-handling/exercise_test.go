package koan

import (
	"errors"
	"strings"
	"testing"
)

func TestProcess(t *testing.T) {
	cause := errors.New("invalid quantity")
	calls := 0
	err := Process(strings.NewReader("ok\nbad\nnever"), func(s string) error {
		calls++
		if s == "bad" {
			return cause
		}
		return nil
	})
	equal(t, calls, 2)
	if !errors.Is(err, cause) {
		t.Fatal("cause lost")
	}
	equal(t, err.Error(), "line 2: invalid quantity")
	noError(t, Process(strings.NewReader(""), func(string) error { return cause }))
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
