//go:build ignore

package koan

type Rectangle struct{ Width, Height int }

func (r Rectangle) Area() int { return r.Width * r.Height }
