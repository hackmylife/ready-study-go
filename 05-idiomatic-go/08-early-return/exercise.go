package koan

import (
	"errors"
)

func CanShip(paid bool, stock int) error { // TODO: elseをなくして平坦にする
	if paid {
		if stock > 0 {
			return nil
		} else {
			return errors.New("out of stock")
		}
	} else {
		return errors.New("unpaid")
	}
}
