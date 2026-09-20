//go:build ignore

package koan

import (
	"context"
)

func Generate(ctx context.Context) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := 0; ; n++ {
			select {
			case <-ctx.Done():
				return
			case out <- n:
			}
		}
	}()
	return out
}
