package koan

import (
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Add(n int) { // TODO: 更新を保護する
	c.value += n
}
func (c *Counter) Value() int { return c.value }
