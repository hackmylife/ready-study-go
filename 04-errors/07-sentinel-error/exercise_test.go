package koan

import (
	"errors"
	"reflect"
	"testing"
)

func TestPop(t *testing.T) {
	v, rest, err := Pop([]int{0, 4})
	noError(t, err)
	equal(t, v, 0)
	same(t, rest, []int{4})
	_, _, err = Pop(nil)
	if !errors.Is(err, ErrEmpty) {
		t.Fatalf("want ErrEmpty, got %v", err)
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
func noError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
