package utils

import "strings"

func NormalizeCode(s string) string { return strings.ToUpper(strings.TrimSpace(s)) }
