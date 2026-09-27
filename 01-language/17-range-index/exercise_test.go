package koan

import (
	"reflect"
	"testing"
)

func TestDoubleInPlace(t *testing.T) {
	values := []int{99, 2, -3, 0, 88}
	DoubleInPlace(values[1:4])
	if want := []int{99, 4, -6, 0, 88}; !reflect.DeepEqual(values, want) {
		t.Fatalf("got %v; want %v", values, want)
	}
	DoubleInPlace(nil)
	DoubleInPlace([]int{})
}
