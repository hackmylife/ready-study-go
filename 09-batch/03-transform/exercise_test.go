package koan

import (
	"reflect"
	"testing"
)

func TestNormalize(t *testing.T) {
	in := Order{Name: "  TEA ", Quantity: 3}
	same(t, Normalize(in), Order{Name: "tea", Quantity: 3})
	equal(t, in.Name, "  TEA ")
	same(t, Normalize(Order{}), Order{})
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
func same(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v; want %#v", got, want)
	}
}
