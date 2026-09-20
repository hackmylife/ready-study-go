//go:build ignore

package koan

func SumEvenBeforeNegative(values []int) int {
	total, i := 0, 0
	for {
		if i == len(values) {
			break
		}
		value := values[i]
		i++
		if value < 0 {
			break
		}
		if value%2 != 0 {
			continue
		}
		total += value
	}
	return total
}
