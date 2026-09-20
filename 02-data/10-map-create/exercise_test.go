package koan

import (
	"reflect"
	"testing"
)

func TestIndex(t *testing.T) {
	same(t, Index([]string{"a", "b", "a"}), map[string]int{"a": 2, "b": 1})
	out := Index(nil)
	if out == nil {
		t.Fatal("map must be writable")
	}
	out["x"] = 1
}
func same(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v; want %#v", got, want)
	}
}
