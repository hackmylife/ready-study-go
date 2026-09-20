package koan

import (
	"testing"
	"time"
)

func TestIssued(t *testing.T) {
	calls := 0
	got := IssuedAt(func() time.Time { calls++; return time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC) })
	equal(t, got, "2026-01-02T03:04:05Z")
	equal(t, calls, 1)
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
