package koan

import (
	"time"
)

func IssuedAt(now func() time.Time) string { // TODO: 渡された時計を使う
	return time.Now().Format(time.RFC3339)
}
