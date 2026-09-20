//go:build ignore

package koan

import (
	"errors"
	"strings"
)

type Order struct {
	Name     string
	Quantity int
}

func Validate(order Order) error {
	if strings.TrimSpace(order.Name) == "" {
		return errors.New("name is required")
	}
	if order.Quantity <= 0 {
		return errors.New("quantity must be positive")
	}
	return nil
}
