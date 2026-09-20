//go:build ignore

package koan

import (
	"errors"
	"time"
)

type RetryError struct{ Delay time.Duration }

func (e *RetryError) Error() string { return "retry later" }
func RetryDelay(err error) (time.Duration, bool) {
	var target *RetryError
	if errors.As(err, &target) {
		return target.Delay, true
	}
	return 0, false
}
