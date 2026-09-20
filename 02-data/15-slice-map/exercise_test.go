package koan

import (
	"reflect"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	values := []int{2, -3, 2}
	var calls []int
	got := Map(values, func(v int) string { calls = append(calls, v); return strconv.Itoa(v) })
	if want := []string{"2", "-3", "2"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v; want %v", got, want)
	}
	if !reflect.DeepEqual(calls, values) {
		t.Fatalf("callback order/count: %v", calls)
	}
	if !reflect.DeepEqual(values, []int{2, -3, 2}) {
		t.Fatal("input changed")
	}
	identity := Map(values, func(v int) int { return v })
	identity[0] = 100
	if values[0] != 2 {
		t.Fatal("output shares input storage")
	}
}
func TestMapEmpty(t *testing.T) {
	f := func(int) string { t.Fatal("callback on empty input"); return "" }
	if Map([]int(nil), f) != nil {
		t.Fatal("nil input must return nil")
	}
	if got := Map([]int{}, f); got == nil || len(got) != 0 {
		t.Fatalf("empty input: %#v", got)
	}
}
