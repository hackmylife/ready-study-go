//go:build ignore

package koan

func DoubleInPlace(values []int) {
	for i := range values {
		values[i] *= 2
	}
}
