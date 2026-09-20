//go:build ignore

package koan

import (
	"testing"
)

func assertBool(t *testing.T, s string, want bool) {
	t.Helper()
	got, err := ParseBool(s)
	if err != nil || got != want {
		t.Fatalf("ParseBool(%q) = %v,%v; want %v,nil", s, got, err, want)
	}
}
func TestParseBool(t *testing.T) {
	assertBool(t, "true", true)
	assertBool(t, "false", false)
	assertBool(t, "1", true)
	_, err := ParseBool("invalid")
	wantError(t, err)
}
func wantError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("expected an error")
	}
}
