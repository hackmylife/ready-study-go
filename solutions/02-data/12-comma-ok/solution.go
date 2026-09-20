//go:build ignore

package koan

func Lookup(m map[string]int, key string) (int, bool) { value, ok := m[key]; return value, ok }
