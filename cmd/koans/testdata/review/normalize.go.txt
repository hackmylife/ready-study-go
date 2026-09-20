//go:build ignore

package koan

import (
	"strings"
)

func NormalizeCode(s string) string { return strings.ToUpper(strings.TrimSpace(s)) }
