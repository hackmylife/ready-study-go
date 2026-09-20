package koan

import (
	"sync"
)

func Count(n int) int { // TODO: 共有変数の競合を修正する
	total := 0
	var wg sync.WaitGroup
	for range n {
		wg.Add(1)
		go func() { defer wg.Done(); total++ }()
	}
	wg.Wait()
	return total
}
