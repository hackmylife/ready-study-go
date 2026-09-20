//go:build ignore

package koan

import (
	"errors"
)

func CanShip(paid bool, stock int) error {
	if !paid {
		return errors.New("unpaid")
	}
	if stock <= 0 {
		return errors.New("out of stock")
	}
	return nil
}
