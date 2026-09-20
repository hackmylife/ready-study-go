//go:build ignore

package koan

func FlatMap[T, U any](values []T, f func(T) []U) []U {
	var out []U
	for _, value := range values {
		out = append(out, f(value)...)
	}
	return out
}
