package koan

import (
	"testing"
	"time"
)

func TestExpiresAt(t *testing.T) {
	start := time.Date(2026, 1, 1, 23, 50, 0, 0, time.UTC)
	equal(t, ExpiresAt(start, 20*time.Minute), time.Date(2026, 1, 2, 0, 10, 0, 0, time.UTC))
	equal(t, ExpiresAt(start, 0), start)
	equal(t, ExpiresAt(start, -time.Minute), start.Add(-time.Minute))
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
