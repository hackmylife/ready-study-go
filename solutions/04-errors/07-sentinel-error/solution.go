//go:build ignore

package koan

import (
	"errors"
)

var ErrEmpty = errors.New("empty queue")

func Pop(values []int) (int, []int, error) {
	if len(values) == 0 {
		return 0, nil, ErrEmpty
	}
	return values[0], values[1:], nil
}
