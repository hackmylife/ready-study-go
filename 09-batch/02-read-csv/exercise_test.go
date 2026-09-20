package koan

import (
	"reflect"
	"strings"
	"testing"
)

func TestReadOrders(t *testing.T) {
	got, err := ReadOrders(strings.NewReader("name,quantity\n\"tea, green\",2\ncoffee,-1\n"))
	noError(t, err)
	same(t, got, []Order{{"tea, green", 2}, {"coffee", -1}})
	for _, s := range []string{"x,y\na,1", "name,quantity\na,x", "name,quantity\na"} {
		_, err := ReadOrders(strings.NewReader(s))
		wantError(t, err)
	}
}
func same(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v; want %#v", got, want)
	}
}
func noError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
func wantError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("expected an error")
	}
}
