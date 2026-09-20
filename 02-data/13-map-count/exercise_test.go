package koan

import (
	"reflect"
	"testing"
)

func TestCounts(t *testing.T) {
	same(t, Counts([]string{"Go", "go", "Go"}), map[string]int{"Go": 2, "go": 1})
	same(t, Counts(nil), map[string]int{})
}
func same(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v; want %#v", got, want)
	}
}
