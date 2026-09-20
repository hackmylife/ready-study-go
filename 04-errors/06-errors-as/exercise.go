package koan

import (
	"time"
)

type RetryError struct{ Delay time.Duration }

func (e *RetryError) Error() string { return "retry later" }
func RetryDelay(err error) (time.Duration, bool) { // TODO: 型付きエラーを探す
	return 0, false
}
