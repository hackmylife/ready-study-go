//go:build ignore

package koan

import (
	"context"
	"time"
)

func Poll(ctx context.Context, interval time.Duration, ready func() bool) error {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			if ready() {
				return nil
			}
		}
	}
}
