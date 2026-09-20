package koan

import (
	"io"
)

type Order struct {
	Name     string
	Quantity int
}

func ReadOrders(r io.Reader) ([]Order, error) { // TODO: CSVからレコードを作る
	return nil, nil
}
