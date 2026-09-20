package koan

import (
	"context"
)

func Process(ctx context.Context, rows []string, handle func(context.Context, string) error) error { // TODO: 各レコードでキャンセルを見る
	return nil
}
