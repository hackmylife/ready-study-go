package koan

import (
	"context"
	"net/http"
)

func Handler(load func(context.Context) (string, error)) http.Handler { // TODO: requestの寿命を伝える
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(503) })
}
