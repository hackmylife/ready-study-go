package koan

import (
	"strconv"
	"testing"
)

func TestParsePort(t *testing.T) {
	for _, s := range []string{"1", "8080", "65535"} {
		n, err := ParsePort(s)
		noError(t, err)
		equal(t, strconv.Itoa(n), s)
	}
	for _, s := range []string{"0", "65536", "x", " 80", "", "99999999999999999999999999"} {
		_, err := ParsePort(s)
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
