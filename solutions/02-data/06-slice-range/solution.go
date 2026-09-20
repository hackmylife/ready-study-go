//go:build ignore

package koan

func Double(values []int) []int {
	if values == nil {
		return nil
	}
	out := make([]int, len(values))
	for i, v := range values {
		out[i] = 2 * v
	}
	return out
}
