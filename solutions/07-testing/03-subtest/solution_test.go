//go:build ignore

package koan

import (
	"testing"
)

func TestSlug(t *testing.T) {
	for _, tt := range []struct{ name, in, want string }{{"trim", " go ", "go"}, {"lower", "GO", "go"}, {"empty", "", ""}} {
		t.Run(tt.name, func(t *testing.T) { equal(t, Slug(tt.in), tt.want) })
	}
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
