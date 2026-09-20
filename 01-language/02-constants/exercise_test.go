package koan

import (
	"testing"
)

func TestMinutes(t *testing.T) {
	equal(t, MinutesPerHour, 60)
	equal(t, HoursPerDay, 24)
	equal(t, MinutesInDays(0), 0)
	equal(t, MinutesInDays(3), 4320)
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
