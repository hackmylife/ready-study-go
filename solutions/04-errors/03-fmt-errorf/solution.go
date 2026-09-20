//go:build ignore

package koan

import (
	"fmt"
)

func ValidateAge(age int) error {
	if age < 0 || age > 150 {
		return fmt.Errorf("age out of range: %d", age)
	}
	return nil
}
