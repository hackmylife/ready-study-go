package koan

import (
	"reflect"
	"testing"
)

func TestZeros(t *testing.T) { same(t, Zeros(3), []int{0, 0, 0}); same(t, Zeros(0), []int{}) }
func same(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v; want %#v", got, want)
	}
}
