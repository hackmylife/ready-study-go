//go:build ignore

package koan

func Reduce[T, U any](values []T, initial U, f func(U, T) U) U {
	for _, value := range values {
		initial = f(initial, value)
	}
	return initial
}
