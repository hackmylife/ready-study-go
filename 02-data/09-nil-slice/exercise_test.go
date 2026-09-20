package koan

import (
	"testing"
)

func TestJSONList(t *testing.T) {
	for _, v := range [][]string{nil, {}} {
		b, err := JSONList(v)
		noError(t, err)
		equal(t, string(b), "[]")
	}
	b, err := JSONList([]string{"a"})
	noError(t, err)
	equal(t, string(b), `["a"]`)
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
