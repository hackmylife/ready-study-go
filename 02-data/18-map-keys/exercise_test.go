package koan

import (
	"reflect"
	"testing"
)

func TestSortedKeys(t *testing.T) {
	values := map[string]int{"pear": 1, "apple": 9, "banana": -2, "": 7}
	want := []string{"", "apple", "banana", "pear"}
	for range 20 {
		if got := SortedKeys(values); !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	}
	if !reflect.DeepEqual(values, map[string]int{"pear": 1, "apple": 9, "banana": -2, "": 7}) {
		t.Fatal("map changed")
	}
	for _, empty := range []map[string]int{nil, {}} {
		if got := SortedKeys(empty); got != nil {
			t.Fatalf("empty result: %#v", got)
		}
	}
}
