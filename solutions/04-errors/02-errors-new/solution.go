//go:build ignore

package koan

import (
	"errors"
)

func ValidateName(name string) error {
	if name == "" {
		return errors.New("name is required")
	}
	return nil
}
