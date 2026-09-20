package koan

import (
	"context"
)

type NameLookup interface {
	Name(context.Context, string) (string, error)
}

func Greeting(ctx context.Context, lookup NameLookup, id string) (string, error) { // TODO: 利用側のinterfaceで取得する
	return "", nil
}
