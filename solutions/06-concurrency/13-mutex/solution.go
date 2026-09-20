//go:build ignore

package koan

import (
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Add(n int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value += n
}
func (c *Counter) Value() int { c.mu.Lock(); defer c.mu.Unlock(); return c.value }
