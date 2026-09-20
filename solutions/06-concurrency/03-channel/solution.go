//go:build ignore

package koan

func Sum(in <-chan int) int {
	sum := 0
	for n := range in {
		sum += n
	}
	return sum
}
