package koan

import (
	"strings"
)

func Slug(s string) string { return strings.ToLower(strings.TrimSpace(s)) }
