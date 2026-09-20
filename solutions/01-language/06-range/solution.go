//go:build ignore

package koan

func CountPositive(values []int) int {
	count := 0
	for _, v := range values {
		if v > 0 {
			count++
		}
	}
	return count
}
