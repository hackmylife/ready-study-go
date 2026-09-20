package koan

import (
	"context"
	"io"
)

func Run(ctx context.Context, workers int, r io.Reader, w io.Writer, work func(context.Context, int) (int, error)) error { // TODO: 数値行を並行変換して順に出力する
	return nil
}
