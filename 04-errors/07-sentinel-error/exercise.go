package koan

import (
	"errors"
)

var ErrEmpty = errors.New("empty queue")

func Pop(values []int) (int, []int, error) { // TODO: 空のキューを判定可能にする
	return 0, nil, nil
}
