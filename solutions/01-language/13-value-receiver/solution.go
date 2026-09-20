//go:build ignore

package koan

type Point struct{ X, Y int }

func (p Point) Moved(dx, dy int) Point { p.X += dx; p.Y += dy; return p }
