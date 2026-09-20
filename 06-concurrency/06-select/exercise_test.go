package koan

import (
	"context"
	"errors"
	"io"
	"testing"
)

func TestReceive(t *testing.T) {
	in := make(chan int, 1)
	in <- 7
	v, err := Receive(context.Background(), in)
	noError(t, err)
	equal(t, v, 7)
	close(in)
	_, err = Receive(context.Background(), in)
	if !errors.Is(err, io.EOF) {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err = Receive(ctx, make(chan int))
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
func noError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
