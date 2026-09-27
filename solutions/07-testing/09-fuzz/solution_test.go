//go:build ignore

package koan

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverse(f *testing.F) {
	for _, seed := range []string{"", "a", "Go言語", "hello, 世界"} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, s string) {
		if !utf8.ValidString(s) {
			t.Skip("Reverseの入力は正しいUTF-8に限る")
		}
		got := Reverse(s)
		if !utf8.ValidString(got) {
			t.Fatalf("Reverse(%q) = %q; invalid UTF-8", s, got)
		}
		if back := Reverse(got); back != s {
			t.Fatalf("Reverse(Reverse(%q)) = %q", s, back)
		}
		first, _ := utf8.DecodeRuneInString(got)
		last, _ := utf8.DecodeLastRuneInString(s)
		if s != "" && first != last {
			t.Fatalf("Reverse(%q) starts with %q; want %q", s, first, last)
		}
	})
}
