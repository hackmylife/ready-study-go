package koan

type Counter struct{ n int }

func (c *Counter) Add(n int) { // TODO: 値を加算する
}
func (c *Counter) Value() int { return c.n }
