package koan

import (
	"reflect"
	"testing"
)

func TestPositive(t *testing.T) {
	in := []int{-1, 3, 0, 2, 3}
	same(t, Positive(in), []int{3, 2, 3})
	same(t, in, []int{-1, 3, 0, 2, 3})
	same(t, Positive([]int{0, -1}), []int(nil))
}
func same(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v; want %#v", got, want)
	}
}
