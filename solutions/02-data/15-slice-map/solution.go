//go:build ignore

package koan

func Map[T, U any](values []T, f func(T) U) []U {
	if values == nil {
		return nil
	}
	out := make([]U, len(values))
	for i, value := range values {
		out[i] = f(value)
	}
	return out
}
