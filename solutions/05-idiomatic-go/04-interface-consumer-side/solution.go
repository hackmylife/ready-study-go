//go:build ignore

package koan

import (
	"context"
)

type NameLookup interface {
	Name(context.Context, string) (string, error)
}

func Greeting(ctx context.Context, lookup NameLookup, id string) (string, error) {
	name, err := lookup.Name(ctx, id)
	if err != nil {
		return "", err
	}
	return "Hello, " + name, nil
}
