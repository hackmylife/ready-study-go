package koan

import (
	"context"
)

func Receive(ctx context.Context, in <-chan int) (int, error) { // TODO: 通信と終了を同時に待つ
	return 0, nil
}
