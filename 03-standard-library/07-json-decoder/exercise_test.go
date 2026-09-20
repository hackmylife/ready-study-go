package koan

import (
	"strings"
	"testing"
)

func TestDecode(t *testing.T) {
	req, err := Decode(strings.NewReader(`{"name":"Aki"} `))
	noError(t, err)
	equal(t, req.Name, "Aki")
	for _, s := range []string{"", `{"other":1}`, `{"name":"x"} {}`, `{"name":"x"} garbage`} {
		_, err := Decode(strings.NewReader(s))
		wantError(t, err)
	}
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
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
