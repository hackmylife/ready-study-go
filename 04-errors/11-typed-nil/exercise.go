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

func ValidateAge(age int) error { // TODO: 成功時にerr != nilとならないようにする
	return checkAge(age)
}
