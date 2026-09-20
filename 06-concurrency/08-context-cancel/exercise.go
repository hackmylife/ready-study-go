package koan

import (
	"context"
)

func Generate(ctx context.Context) <-chan int { // TODO: 送信待ちにもcancelを効かせる
	out := make(chan int)
	close(out)
	return out
}
