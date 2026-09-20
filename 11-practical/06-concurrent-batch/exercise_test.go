package koan

import (
	"bytes"
	"context"
	"errors"
	"io"
	"strings"
	"sync/atomic"
	"testing"
)

func TestConcurrentBatch(t *testing.T) {
	var b bytes.Buffer
	var active, peak atomic.Int64
	noError(t, Run(context.Background(), 2, strings.NewReader("3\n1\n2"), &b, func(ctx context.Context, v int) (int, error) {
		n := active.Add(1)
		defer active.Add(-1)
		for old := peak.Load(); n > old && !peak.CompareAndSwap(old, n); old = peak.Load() {
		}
		return v * v, nil
	}))
	equal(t, b.String(), "9\n1\n4\n")
	if peak.Load() > 2 {
		t.Fatal("unbounded work")
	}
	cause := errors.New("failed")
	b.Reset()
	err := Run(context.Background(), 2, strings.NewReader("1\n2"), &b, func(context.Context, int) (int, error) { return 0, cause })
	if !errors.Is(err, cause) {
		t.Fatal(err)
	}
	equal(t, b.Len(), 0)
	wantError(t, Run(context.Background(), 2, strings.NewReader("x"), io.Discard, nil))
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if !errors.Is(Run(ctx, 2, strings.NewReader("1"), io.Discard, nil), context.Canceled) {
		t.Fatal("cancel lost")
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
func wantError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("expected an error")
	}
}
