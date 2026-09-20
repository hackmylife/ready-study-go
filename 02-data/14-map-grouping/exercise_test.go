package koan

import (
	"reflect"
	"testing"
)

func TestGroup(t *testing.T) {
	same(t, Group([]Member{{"A", "red"}, {"B", "blue"}, {"C", "red"}}), map[string][]string{"red": {"A", "C"}, "blue": {"B"}})
	same(t, Group(nil), map[string][]string{})
}
func same(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v; want %#v", got, want)
	}
}
