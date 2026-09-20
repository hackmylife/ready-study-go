//go:build ignore

package koan

import (
	"strings"
)

type Order struct {
	Name     string
	Quantity int
}

func Normalize(order Order) Order {
	order.Name = strings.ToLower(strings.TrimSpace(order.Name))
	return order
}
