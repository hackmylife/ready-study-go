//go:build ignore

package koan

import (
	"errors"
	"time"
)

type Client struct{ Timeout time.Duration }
type Option func(*Client) error

func WithTimeout(d time.Duration) Option {
	return func(c *Client) error {
		if d <= 0 {
			return errors.New("timeout must be positive")
		}
		c.Timeout = d
		return nil
	}
}
func NewClient(options ...Option) (*Client, error) {
	c := &Client{Timeout: 5 * time.Second}
	for _, option := range options {
		if err := option(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}
