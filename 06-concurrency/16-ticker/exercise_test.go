package koan

import (
	"context"
	"errors"
	"testing"
	"testing/synctest"
	"time"
)

func TestPollReturnsWhenReady(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		calls := 0
		start := time.Now()
		err := Poll(t.Context(), time.Second, func() bool {
			calls++
			return calls == 3
		})
		if err != nil {
			t.Fatal(err)
		}
		equal(t, calls, 3)
		equal(t, time.Since(start), 3*time.Second)
	})
}

func TestPollStopsWhenContextIsDone(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithTimeout(t.Context(), 2500*time.Millisecond)
		defer cancel()
		calls := 0
		start := time.Now()
		err := Poll(ctx, time.Second, func() bool {
			calls++
			return false
		})
		if !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("got %v; want context.DeadlineExceeded", err)
		}
		equal(t, calls, 2)
		equal(t, time.Since(start), 2500*time.Millisecond)
	})
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
