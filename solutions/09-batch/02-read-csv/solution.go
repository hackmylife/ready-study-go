//go:build ignore

package koan

import (
	"encoding/csv"
	"errors"
	"io"
	"slices"
	"strconv"
)

type Order struct {
	Name     string
	Quantity int
}

func ReadOrders(r io.Reader) ([]Order, error) {
	reader := csv.NewReader(r)
	header, err := reader.Read()
	if err != nil {
		return nil, err
	}
	if !slices.Equal(header, []string{"name", "quantity"}) {
		return nil, errors.New("invalid header")
	}
	var out []Order
	for {
		row, err := reader.Read()
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
		n, err := strconv.Atoi(row[1])
		if err != nil {
			return nil, err
		}
		out = append(out, Order{Name: row[0], Quantity: n})
	}
}
