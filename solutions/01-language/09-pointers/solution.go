//go:build ignore

package koan

func Increment(value *int) bool {
	if value == nil {
		return false
	}
	*value += 1
	return true
}
