//go:build ignore

package koan

import (
	"errors"
	"net/http"
	"strconv"
)

func Limit(r *http.Request) (int, error) {
	values, ok := r.URL.Query()["limit"]
	if !ok {
		return 20, nil
	}
	if len(values) != 1 {
		return 0, errors.New("one limit is required")
	}
	n, err := strconv.Atoi(values[0])
	if err != nil {
		return 0, err
	}
	if n < 1 || n > 100 {
		return 0, errors.New("limit must be 1..100")
	}
	return n, nil
}
