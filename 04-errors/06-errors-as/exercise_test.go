package koan

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestRetry(t *testing.T) {
	d, ok := RetryDelay(fmt.Errorf("request: %w", &RetryError{Delay: 3 * time.Second}))
	equal(t, d, 3*time.Second)
	equal(t, ok, true)
	d, ok = RetryDelay(errors.New("other"))
	equal(t, d, time.Duration(0))
	equal(t, ok, false)
	_, ok = RetryDelay(nil)
	equal(t, ok, false)
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
