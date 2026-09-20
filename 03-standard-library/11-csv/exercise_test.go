package koan

import (
	"reflect"
	"strings"
	"testing"
)

func TestCSV(t *testing.T) {
	got, err := ReadCSV(strings.NewReader("name,note\nAki,\"hello, Go\"\n"))
	noError(t, err)
	same(t, got, [][]string{{"name", "note"}, {"Aki", "hello, Go"}})
	_, err = ReadCSV(strings.NewReader("a,b\nonly-one\n"))
	wantError(t, err)
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
