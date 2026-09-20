//go:build ignore

package koan

func ApplyTwice(value int, f func(int) int) int { return f(f(value)) }
