package koan

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestMissing(t *testing.T) {
	equal(t, Missing(os.ErrNotExist), true)
	equal(t, Missing(fmt.Errorf("load: %w", os.ErrNotExist)), true)
	equal(t, Missing(nil), false)
	equal(t, Missing(errors.New(os.ErrNotExist.Error())), false)
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
