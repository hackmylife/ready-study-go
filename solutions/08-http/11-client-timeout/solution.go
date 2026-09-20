//go:build ignore

package koan

import (
	"errors"
	"net/http"
	"time"
)

func NewClient(timeout time.Duration) (*http.Client, error) {
	if timeout <= 0 {
		return nil, errors.New("timeout must be positive")
	}
	return &http.Client{Timeout: timeout}, nil
}
