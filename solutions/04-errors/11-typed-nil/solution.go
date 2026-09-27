//go:build ignore

package koan

import (
	"fmt"
)

type ValidationError struct{ Field string }

func (e *ValidationError) Error() string { return fmt.Sprintf("invalid %s", e.Field) }

func checkAge(age int) *ValidationError {
	if age < 0 {
		return &ValidationError{Field: "age"}
	}
	return nil
}

func ValidateAge(age int) error {
	if err := checkAge(age); err != nil {
		return err
	}
	return nil
}
