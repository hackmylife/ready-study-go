//go:build ignore

package koan

import (
	"errors"
	"testing"
)

func TestLoad(t *testing.T) {
	noError(t, Load(true))
	if !errors.Is(Load(false), ErrMissing) {
		t.Fatalf("expected ErrMissing, got %v", Load(false))
	}
}
func noError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
