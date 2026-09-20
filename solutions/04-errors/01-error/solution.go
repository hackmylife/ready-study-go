//go:build ignore

package koan

import (
	"errors"
	"math"
)

func Sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, errors.New("negative input")
	}
	return math.Sqrt(n), nil
}
