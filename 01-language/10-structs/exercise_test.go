package koan

import (
	"reflect"
	"testing"
)

func TestUser(t *testing.T) {
	same(t, NewUser("Aki", 21), User{Name: "Aki", Age: 21})
	same(t, NewUser("", 0), User{})
}
func same(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v; want %#v", got, want)
	}
}
