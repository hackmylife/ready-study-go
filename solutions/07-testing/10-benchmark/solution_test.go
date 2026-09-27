//go:build ignore

package koan

import (
	"testing"
)

var words = []string{"alpha", "beta", "gamma", "delta", "epsilon", "zeta", "eta", "theta"}

func TestJoin(t *testing.T) {
	for _, tt := range []struct {
		name  string
		parts []string
		want  string
	}{
		{"many", []string{"a", "b", "c"}, "a, b, c"},
		{"one", []string{"a"}, "a"},
		{"empty", nil, ""},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := Join(tt.parts, ", "); got != tt.want {
				t.Fatalf("Join(%q) = %q; want %q", tt.parts, got, tt.want)
			}
		})
	}
}

func TestJoinAllocs(t *testing.T) {
	allocs := testing.AllocsPerRun(100, func() { Join(words, ", ") })
	if allocs > 1 {
		t.Fatalf("Join allocated %v times per call; want at most 1", allocs)
	}
}

func BenchmarkJoin(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		Join(words, ", ")
	}
}
