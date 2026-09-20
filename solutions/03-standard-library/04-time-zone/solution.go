//go:build ignore

package koan

import (
	"time"
)

func InTokyo(t time.Time) (time.Time, error) {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return time.Time{}, err
	}
	return t.In(loc), nil
}
