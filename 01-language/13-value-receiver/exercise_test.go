package koan

import (
	"reflect"
	"testing"
)

func TestMoved(t *testing.T) {
	p := Point{X: 2, Y: 5}
	same(t, p.Moved(3, -2), Point{X: 5, Y: 3})
	same(t, p, Point{X: 2, Y: 5})
}
func same(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v; want %#v", got, want)
	}
}
