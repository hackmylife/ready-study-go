//go:build ignore

package koan

import (
	"context"
)

func Fetch(ctx context.Context, load func(context.Context) (string, error)) (string, error) {
	return load(ctx)
}
