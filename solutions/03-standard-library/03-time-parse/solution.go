//go:build ignore

package koan

import (
	"time"
)

func ParseDate(s string) (time.Time, error) { return time.Parse("2006-01-02", s) }
