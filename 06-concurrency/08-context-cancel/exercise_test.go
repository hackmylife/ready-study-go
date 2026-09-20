package koan

import (
	"context"
	"testing"
	"time"
)

func TestGenerate(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	out := Generate(ctx)
	for want := range 3 {
		select {
		case got, ok := <-out:
			equal(t, ok, true)
			equal(t, got, want)
		case <-time.After(time.Second):
			t.Fatal("no value")
		}
	}
	cancel()
	deadline := time.NewTimer(time.Second)
	defer deadline.Stop()
	for {
		select {
		case _, ok := <-out:
			if !ok {
				return
			}
		case <-deadline.C:
			t.Fatal("generator did not stop")
		}
	}
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
