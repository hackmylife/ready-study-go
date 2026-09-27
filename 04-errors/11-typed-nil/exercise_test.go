package koan

import (
	"errors"
	"testing"
)

func TestValidateAge(t *testing.T) {
	for _, age := range []int{20, 0} {
		if err := ValidateAge(age); err != nil {
			t.Fatalf("ValidateAge(%d) = %#v; want nil", age, err)
		}
	}
	var target *ValidationError
	if err := ValidateAge(-1); !errors.As(err, &target) || target.Field != "age" {
		t.Fatalf("ValidateAge(-1) = %#v; want *ValidationError{Field: \"age\"}", err)
	}
}
