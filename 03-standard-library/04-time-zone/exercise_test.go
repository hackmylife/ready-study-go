package koan

import (
	"testing"
	"time"
)

func TestTokyo(t *testing.T) {
	in := time.Date(2026, 1, 1, 16, 0, 0, 0, time.UTC)
	out, err := InTokyo(in)
	noError(t, err)
	equal(t, out.Equal(in), true)
	equal(t, out.Format("2006-01-02 15:04 -0700"), "2026-01-02 01:00 +0900")
	equal(t, out.Location().String(), "Asia/Tokyo")
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
