//go:build ignore

package koan

import (
	"time"
)

func ExpiresAt(start time.Time, ttl time.Duration) time.Time { return start.Add(ttl) }
