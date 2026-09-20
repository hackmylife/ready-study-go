//go:build ignore

package koan

import (
	"context"
	"io"
)

func Receive(ctx context.Context, in <-chan int) (int, error) {
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case v, ok := <-in:
		if !ok {
			return 0, io.EOF
		}
		return v, nil
	}
}
