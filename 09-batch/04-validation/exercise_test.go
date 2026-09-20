package koan

import (
	"testing"
)

func TestValidate(t *testing.T) {
	noError(t, Validate(Order{"tea", 1}))
	for _, v := range []Order{{"", 1}, {" ", 1}, {"tea", 0}, {"tea", -2}} {
		wantError(t, Validate(v))
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
