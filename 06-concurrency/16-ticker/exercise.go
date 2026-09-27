package koan

import (
	"context"
	"time"
)

func Poll(ctx context.Context, interval time.Duration, ready func() bool) error { // TODO: intervalごとにreadyを確認する
	return nil
}
