//go:build ignore

package koan

type Counter struct{ n int }

func (c *Counter) Add(n int)  { c.n += n }
func (c *Counter) Value() int { return c.n }
