//go:build ignore

package koan

func Positive(values []int) []int {
	var out []int
	for _, v := range values {
		if v > 0 {
			out = append(out, v)
		}
	}
	return out
}
