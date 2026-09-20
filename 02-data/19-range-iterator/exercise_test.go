package koan

import (
	"reflect"
	"slices"
	"testing"
)

func TestTake(t *testing.T) {
	for _, tt := range []struct {
		values []int
		n      int
		want   []int
	}{
		{[]int{4, 5, 6}, 2, []int{4, 5}}, {[]int{4}, 3, []int{4}}, {nil, 3, nil}, {[]int{4}, 0, nil}, {[]int{4}, -1, nil},
	} {
		if got := Take(slices.Values(tt.values), tt.n); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("input=%v n=%d: got %v; want %v", tt.values, tt.n, got, tt.want)
		}
	}
}
func TestTakeStopsWithoutExtraRead(t *testing.T) {
	visited, stopped := 0, false
	seq := func(yield func(int) bool) {
		defer func() { stopped = true }()
		for value := 10; ; value++ {
			visited++
			if visited > 2 {
				t.Fatal("requested an extra element")
			}
			if !yield(value) {
				return
			}
		}
	}
	got := Take(seq, 2)
	if !reflect.DeepEqual(got, []int{10, 11}) || visited != 2 || !stopped {
		t.Fatalf("got=%v visited=%d stopped=%v", got, visited, stopped)
	}
	for _, n := range []int{0, -1} {
		Take(func(func(int) bool) { t.Fatal("started iterator for nonpositive n") }, n)
	}
}
