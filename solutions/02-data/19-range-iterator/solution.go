//go:build ignore

package koan

import "iter"

func Take(seq iter.Seq[int], n int) []int {
	if n <= 0 {
		return nil
	}
	var out []int
	for value := range seq {
		out = append(out, value)
		if len(out) == n {
			break
		}
	}
	return out
}
