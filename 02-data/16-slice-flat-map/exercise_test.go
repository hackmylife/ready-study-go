package koan

import (
	"reflect"
	"strings"
	"testing"
)

func TestFlatMap(t *testing.T) {
	input := []string{"Go is", "", "fun Go"}
	var calls []string
	got := FlatMap(input, func(s string) []string { calls = append(calls, s); return strings.Fields(s) })
	if want := []string{"Go", "is", "fun", "Go"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v; want %v", got, want)
	}
	if !reflect.DeepEqual(calls, input) {
		t.Fatalf("callback order/count: %v", calls)
	}
	if !reflect.DeepEqual(input, []string{"Go is", "", "fun Go"}) {
		t.Fatal("input changed")
	}
	if got := FlatMap([]string{"", "  "}, strings.Fields); got != nil {
		t.Fatalf("empty result: %#v", got)
	}
	for _, empty := range [][]string{nil, {}} {
		if got := FlatMap(empty, func(string) []int { t.Fatal("callback on empty input"); return nil }); got != nil {
			t.Fatalf("empty input: %#v", got)
		}
	}
}
func TestFlatMapOwnsResult(t *testing.T) {
	storage := []int{1, 2, 99, 99}
	parts := [][]int{storage[:2], nil, {3}}
	got := FlatMap(parts, func(v []int) []int { return v })
	if !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Fatalf("got %v", got)
	}
	if storage[2] != 99 {
		t.Fatal("append changed callback storage")
	}
	got[0] = 10
	if storage[0] != 1 {
		t.Fatal("output shares callback storage")
	}
}
