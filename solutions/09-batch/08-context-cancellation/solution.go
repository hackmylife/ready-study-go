//go:build ignore

package koan

import (
	"context"
)

func Process(ctx context.Context, rows []string, handle func(context.Context, string) error) error {
	for _, row := range rows {
		if err := ctx.Err(); err != nil {
			return err
		}
		if err := handle(ctx, row); err != nil {
			return err
		}
	}
	return ctx.Err()
}
