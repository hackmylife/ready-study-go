package koan

import (
	"reflect"
	"testing"
)

func TestAppend(t *testing.T) {
	same(t, AppendValue([]string{"a", "b"}, "c"), []string{"a", "b", "c"})
	same(t, AppendValue(nil, "x"), []string{"x"})
}
func same(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v; want %#v", got, want)
	}
}
