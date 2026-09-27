//go:build ignore

package koan

import (
	"context"
	"errors"
	"testing"
	"testing/synctest"
	"time"
)

var errTemporary = errors.New("temporary")

func TestRetrySucceedsAfterDelay(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		calls := 0
		start := time.Now()
		err := Retry(t.Context(), 5, time.Second, func() error {
			calls++
			if calls < 3 {
				return errTemporary
			}
			return nil
		})
		if err != nil {
			t.Fatal(err)
		}
		equal(t, calls, 3)
		equal(t, time.Since(start), 2*time.Second)
	})
}

func TestRetryReturnsLastError(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		errs := []error{errors.New("first"), errors.New("second"), errors.New("third"), errors.New("extra")}
		calls := 0
		err := Retry(t.Context(), 3, time.Second, func() error {
			calls++
			return errs[calls-1]
		})
		if !errors.Is(err, errs[2]) {
			t.Fatalf("got %v; want third", err)
		}
		equal(t, calls, 3)
	})
}

func TestRetryStopsWhenContextIsDone(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithTimeout(t.Context(), 1500*time.Millisecond)
		defer cancel()
		calls := 0
		start := time.Now()
		err := Retry(ctx, 5, time.Second, func() error {
			calls++
			return errTemporary
		})
		if !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("got %v; want context.DeadlineExceeded", err)
		}
		equal(t, calls, 2)
		equal(t, time.Since(start), 1500*time.Millisecond)
	})
}

func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
