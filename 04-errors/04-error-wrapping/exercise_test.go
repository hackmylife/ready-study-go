package koan

import (
	"errors"
	"testing"
)

func TestLoad(t *testing.T) {
	cause := errors.New("offline")
	_, err := Load("users.json", func(p string) ([]byte, error) { equal(t, p, "users.json"); return nil, cause })
	if !errors.Is(err, cause) {
		t.Fatal("cause lost")
	}
	equal(t, err.Error(), "load users.json: offline")
	b, err := Load("ok", func(string) ([]byte, error) { return []byte("ok"), nil })
	noError(t, err)
	equal(t, string(b), "ok")
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
