//go:build ignore

package koan

import (
	"testing"
)

func TestGrade(t *testing.T) {
	for _, tt := range []struct {
		score int
		want  string
	}{{0, "C"}, {59, "C"}, {60, "B"}, {79, "B"}, {80, "A"}, {100, "A"}} {
		equal(t, Grade(tt.score), tt.want)
	}
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
