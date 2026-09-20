//go:build ignore

package koan

import (
	"context"
	"errors"
	"net/http"
	"time"
)

func Stop(server *http.Server, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		return errors.Join(err, server.Close())
	}
	return nil
}
