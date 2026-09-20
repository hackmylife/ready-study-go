//go:build ignore

package koan

import (
	"time"
)

func IssuedAt(now func() time.Time) string { return now().Format(time.RFC3339) }
