package koan

import (
	"errors"
	"reflect"
	"testing"
)

func TestUse(t *testing.T) {
	for _, cause := range []error{nil, errors.New("failed")} {
		var events []string
		err := Use(func() error { events = append(events, "work"); return cause }, func() { events = append(events, "close") })
		equal(t, err, cause)
		same(t, events, []string{"work", "close"})
	}
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
