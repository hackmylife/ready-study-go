package koan

import (
	"errors"
	"testing"
)

func TestValidate(t *testing.T) {
	err := Validate("")
	var field *FieldError
	if !errors.As(err, &field) {
		t.Fatalf("want FieldError, got %v", err)
	}
	equal(t, field.Field, "email")
	equal(t, field.Reason, "required")
	equal(t, err.Error(), "email: required")
	noError(t, Validate("a@example.test"))
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
