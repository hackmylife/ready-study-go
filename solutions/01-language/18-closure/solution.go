//go:build ignore

package koan

func Counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
