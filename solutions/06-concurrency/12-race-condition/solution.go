//go:build ignore

package koan

import (
	"sync"
	"sync/atomic"
)

func Count(n int) int {
	var total atomic.Int64
	var wg sync.WaitGroup
	for range n {
		wg.Add(1)
		go func() { defer wg.Done(); total.Add(1) }()
	}
	wg.Wait()
	return int(total.Load())
}
