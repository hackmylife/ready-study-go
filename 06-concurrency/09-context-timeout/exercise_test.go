package koan

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestTimeout(t *testing.T) {
	err := WithTimeout(context.Background(), 0, func(ctx context.Context) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			return errors.New("missing deadline")
		}
	})
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatal(err)
	}
	var child context.Context
	noError(t, WithTimeout(context.Background(), time.Hour, func(ctx context.Context) error { child = ctx; return nil }))
	if child.Err() != context.Canceled {
		t.Fatal("child was not canceled after work")
	}
	parent, cancel := context.WithCancel(context.Background())
	cancel()
	err = WithTimeout(parent, time.Hour, func(ctx context.Context) error { return ctx.Err() })
	if !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}
}
func noError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
