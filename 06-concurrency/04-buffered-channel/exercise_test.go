package koan

import (
	"reflect"
	"testing"
)

func TestQueue(t *testing.T) {
	q := Queue([]int{3, 1, 4})
	equal(t, cap(q), 3)
	var got []int
	for n := range q {
		got = append(got, n)
	}
	same(t, got, []int{3, 1, 4})
	_, ok := <-Queue(nil)
	equal(t, ok, false)
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
