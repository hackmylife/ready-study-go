package koan

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

type broken struct{}

func (broken) Read([]byte) (int, error) { return 0, errors.New("read failed") }
func TestLines(t *testing.T) {
	got, err := Lines(strings.NewReader("a\n\nb\r\nlast"))
	noError(t, err)
	same(t, got, []string{"a", "", "b", "last"})
	got, err = Lines(strings.NewReader(""))
	noError(t, err)
	same(t, got, []string(nil))
	_, err = Lines(broken{})
	wantError(t, err)
	_, err = Lines(strings.NewReader(strings.Repeat("x", 128*1024)))
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
