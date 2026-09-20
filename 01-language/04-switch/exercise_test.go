package koan

import (
	"testing"
)

func TestDayKind(t *testing.T) {
	for _, d := range []string{"Mon", "Tue", "Wed", "Thu", "Fri"} {
		equal(t, DayKind(d), "weekday")
	}
	for _, d := range []string{"Sat", "Sun"} {
		equal(t, DayKind(d), "weekend")
	}
	equal(t, DayKind(""), "invalid")
	equal(t, DayKind("sun"), "invalid")
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
