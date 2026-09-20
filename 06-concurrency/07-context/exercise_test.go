package koan

import (
	"context"
	"errors"
	"testing"
)

func TestFetch(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	got, err := Fetch(ctx, func(received context.Context) (string, error) { equal(t, received, ctx); return "", received.Err() })
	equal(t, got, "")
	if !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
