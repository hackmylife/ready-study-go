//go:build ignore

package koan

type Counter struct{ value int }

func (c *Counter) Add(n int)  { c.value += n }
func (c *Counter) Value() int { return c.value }
