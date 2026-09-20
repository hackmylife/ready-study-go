//go:build ignore

package koan

func Index(names []string) map[string]int {
	out := make(map[string]int)
	for i, n := range names {
		out[n] = i
	}
	return out
}
