//go:build ignore

package koan

import (
	"context"
	"errors"
	"sync"
)

func Map(ctx context.Context, workers int, values []int, work func(context.Context, int) (int, error)) ([]int, error) {
	if workers <= 0 {
		return nil, errors.New("workers must be positive")
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	out := make([]int, len(values))
	jobs := make(chan int)
	var wg sync.WaitGroup
	var once sync.Once
	var first error
	for range min(workers, len(values)) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range jobs {
				if ctx.Err() != nil {
					return
				}
				v, err := work(ctx, values[i])
				if err != nil {
					once.Do(func() { first = err; cancel() })
					return
				}
				out[i] = v
			}
		}()
	}
send:
	for i := range values {
		select {
		case <-ctx.Done():
			break send
		case jobs <- i:
		}
	}
	close(jobs)
	wg.Wait()
	if first != nil {
		return nil, first
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return out, nil
}
