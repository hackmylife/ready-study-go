package koan

import (
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	u, err := Decode([]byte(`{"name":"Aki","age":21,"extra":true}`))
	noError(t, err)
	same(t, u, User{Name: "Aki", Age: 21})
	for _, s := range []string{`{`, `{"age":"21"}`} {
		_, err := Decode([]byte(s))
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
