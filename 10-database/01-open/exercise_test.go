package koan

import (
	"context"
	"errors"
	"ready-study-go/internal/dbtest"
	"testing"
)

func TestConnect(t *testing.T) {
	fixture := dbtest.New(t)
	conn, err := Connect(context.Background(), fixture.DSN)
	noError(t, err)
	if conn == nil {
		t.Fatal("nil DB")
	}
	defer conn.Close()
	noError(t, conn.PingContext(context.Background()))
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	conn, err = Connect(ctx, fixture.DSN)
	if conn != nil {
		conn.Close()
	}
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("want canceled, got %v", err)
	}
}
func noError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
