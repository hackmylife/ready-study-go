//go:build ignore

package koan

func Indices(n int) []int {
	var out []int
	for i := range n {
		out = append(out, i)
	}
	return out
}
