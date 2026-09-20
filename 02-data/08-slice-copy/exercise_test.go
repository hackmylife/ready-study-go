package koan

import (
	"reflect"
	"testing"
)

func TestClone(t *testing.T) {
	in := []int{2, 4}
	out := Clone(in)
	same(t, out, in)
	out[0] = 9
	equal(t, in[0], 2)
	same(t, Clone(nil), []int(nil))
	same(t, Clone([]int{}), []int{})
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
