package koan

import (
	"strings"
)

func Join(parts []string, sep string) string {
	if len(parts) == 0 {
		return ""
	}
	size := len(sep) * (len(parts) - 1)
	for _, part := range parts {
		size += len(part)
	}
	var b strings.Builder
	b.Grow(size)
	for i, part := range parts {
		if i > 0 {
			b.WriteString(sep)
		}
		b.WriteString(part)
	}
	return b.String()
}
