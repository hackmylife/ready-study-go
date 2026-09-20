//go:build ignore

package koan

import (
	"strings"
)

func Normalize(s string) string { return strings.Join(strings.Fields(s), " ") }
