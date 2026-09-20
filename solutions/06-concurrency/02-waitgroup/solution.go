//go:build ignore

package koan

import (
	"sync"
)

func RunAll(jobs []func()) {
	var wg sync.WaitGroup
	for _, job := range jobs {
		wg.Add(1)
		go func() { defer wg.Done(); job() }()
	}
	wg.Wait()
}
