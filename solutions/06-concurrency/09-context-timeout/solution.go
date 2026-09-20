//go:build ignore

package koan

import (
	"context"
	"time"
)

func WithTimeout(ctx context.Context, d time.Duration, work func(context.Context) error) error {
	child, cancel := context.WithTimeout(ctx, d)
	defer cancel()
	return work(child)
}
