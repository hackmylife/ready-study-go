package koan

import (
	"reflect"
	"testing"
)

func TestDouble(t *testing.T) {
	in := []int{2, -3, 0}
	same(t, Double(in), []int{4, -6, 0})
	same(t, in, []int{2, -3, 0})
	same(t, Double(nil), []int(nil))
	same(t, Double([]int{}), []int{})
}
func same(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v; want %#v", got, want)
	}
}
