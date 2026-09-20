package koan

import (
	"testing"
	"time"
)

func TestParseDate(t *testing.T) {
	got, err := ParseDate("2024-02-29")
	noError(t, err)
	equal(t, got, time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC))
	for _, s := range []string{"2023-02-29", "2024/02/29", ""} {
		_, err := ParseDate(s)
		wantError(t, err)
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
func wantError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("expected an error")
	}
}
