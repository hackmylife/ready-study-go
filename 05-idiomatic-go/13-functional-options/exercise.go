package koan

import (
	"time"
)

type Client struct{ Timeout time.Duration }
type Option func(*Client) error

func WithTimeout(d time.Duration) Option { // TODO: optionを作る
	return func(*Client) error { return nil }
}
func NewClient(options ...Option) (*Client, error) { return &Client{}, nil }
