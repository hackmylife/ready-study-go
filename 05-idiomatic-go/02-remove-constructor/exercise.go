package koan

type Counter struct{ value int }

func NewCounter() *Counter    { return &Counter{} } // TODO: 不要なconstructorを削除する
func (c *Counter) Add(n int)  { c.value += n }
func (c *Counter) Value() int { return c.value }
