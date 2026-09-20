package koan

import (
	"context"
	"errors"
	"ready-study-go/internal/dbtest"
	"testing"
)

func TestPing(t *testing.T) {
	db := dbtest.Open(t)
	noError(t, Ping(context.Background(), db))
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if !errors.Is(Ping(ctx, db), context.Canceled) {
		t.Fatal("cancellation was lost")
	}
}
func noError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
