package koan

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestEachLine(t *testing.T) {
	var got []string
	noError(t, EachLine(strings.NewReader("a\n\nb"), func(s string) error { got = append(got, s); return nil }))
	same(t, got, []string{"a", "", "b"})
	cause := errors.New("stop")
	calls := 0
	err := EachLine(strings.NewReader("a\nb"), func(string) error { calls++; return cause })
	if !errors.Is(err, cause) {
		t.Fatal(err)
	}
	equal(t, calls, 1)
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
func same(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v; want %#v", got, want)
	}
}
func noError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
