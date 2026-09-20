package koan

import (
	"reflect"
	"strconv"
	"testing"
)

func TestReduce(t *testing.T) {
	values := []int{1, 2, 3}
	var calls []int
	got := Reduce(values, "start", func(acc string, v int) string { calls = append(calls, v); return acc + "/" + strconv.Itoa(v) })
	if got != "start/1/2/3" {
		t.Fatalf("got %q", got)
	}
	if !reflect.DeepEqual(calls, values) {
		t.Fatalf("callback order/count: %v", calls)
	}
	if !reflect.DeepEqual(values, []int{1, 2, 3}) {
		t.Fatal("input changed")
	}
	if got := Reduce(values, 10, func(acc, v int) int { return acc + v }); got != 16 {
		t.Fatalf("sum=%d", got)
	}
	for _, empty := range [][]int{nil, {}} {
		if got := Reduce(empty, 10, func(acc, v int) int { t.Fatal("callback on empty input"); return 0 }); got != 10 {
			t.Fatalf("empty result=%d", got)
		}
	}
}
