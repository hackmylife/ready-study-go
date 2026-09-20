//go:build ignore

package koan

import (
	"maps"
	"slices"
)

func SortedKeys(values map[string]int) []string {
	return slices.Sorted(maps.Keys(values))
}
