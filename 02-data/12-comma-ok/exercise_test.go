package koan

import (
	"testing"
)

func TestLookup(t *testing.T) {
	m := map[string]int{"zero": 0, "one": 1}
	for _, tt := range []struct {
		k  string
		v  int
		ok bool
	}{{"zero", 0, true}, {"one", 1, true}, {"missing", 0, false}} {
		v, ok := Lookup(m, tt.k)
		equal(t, v, tt.v)
		equal(t, ok, tt.ok)
	}
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
