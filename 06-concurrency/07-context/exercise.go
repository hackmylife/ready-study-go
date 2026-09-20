package koan

import (
	"context"
)

func Fetch(ctx context.Context, load func(context.Context) (string, error)) (string, error) { // TODO: contextを伝播する
	return load(context.Background())
}
