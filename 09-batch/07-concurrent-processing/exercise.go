package koan

import (
	"context"
)

func Map(ctx context.Context, workers int, values []int, work func(context.Context, int) (int, error)) ([]int, error) { // TODO: worker数を制限して順序を保つ
	return nil, nil
}
