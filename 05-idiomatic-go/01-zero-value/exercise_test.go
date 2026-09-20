package koan

import (
	"testing"
)

func TestSet(t *testing.T) {
	var s Set
	equal(t, s.Has("Go"), false)
	s.Add("Go")
	s.Add("Go")
	equal(t, s.Has("Go"), true)
	equal(t, s.Has("Rust"), false)
	var other Set
	equal(t, other.Has("Go"), false)
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
