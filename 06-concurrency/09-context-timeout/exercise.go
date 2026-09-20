package koan

import (
	"context"
	"time"
)

func WithTimeout(ctx context.Context, d time.Duration, work func(context.Context) error) error { // TODO: 子contextを作り後始末する
	return work(ctx)
}
