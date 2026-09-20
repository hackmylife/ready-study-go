package koan

import (
	"testing"
)

func TestValidateName(t *testing.T) {
	err := ValidateName("")
	wantError(t, err)
	equal(t, err.Error(), "name is required")
	noError(t, ValidateName("Aki"))
	noError(t, ValidateName(" "))
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
