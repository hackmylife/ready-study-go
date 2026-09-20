package koan

import (
	"context"
	"errors"
	"reflect"
	"sync/atomic"
	"testing"
	"time"
)

func TestMap(t *testing.T) {
	got, err := Map(context.Background(), 2, []int{3, 1, 2}, func(ctx context.Context, v int) (int, error) { return v * 2, nil })
	noError(t, err)
	same(t, got, []int{6, 2, 4})
	got, err = Map(context.Background(), 2, nil, func(context.Context, int) (int, error) { t.Fatal("unexpected job"); return 0, nil })
	noError(t, err)
	same(t, got, []int{})
	_, err = Map(context.Background(), 0, nil, nil)
	wantError(t, err)
}
func TestMapError(t *testing.T) {
	cause := errors.New("job failed")
	_, err := Map(context.Background(), 2, []int{1, 2, 3}, func(context.Context, int) (int, error) { return 0, cause })
	if !errors.Is(err, cause) {
		t.Fatalf("error lost: %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err = Map(ctx, 2, []int{1}, func(context.Context, int) (int, error) { t.Error("called after cancel"); return 0, nil })
	if !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}
}
func TestMapLimit(t *testing.T) {
	var active, peak atomic.Int64
	gate := make(chan struct{})
	started := make(chan struct{}, 4)
	done := make(chan struct{})
	var got []int
	var err error
	go func() {
		got, err = Map(context.Background(), 2, []int{1, 2, 3, 4}, func(ctx context.Context, v int) (int, error) {
			n := active.Add(1)
			defer active.Add(-1)
			for old := peak.Load(); n > old && !peak.CompareAndSwap(old, n); old = peak.Load() {
			}
			started <- struct{}{}
			<-gate
			return v, nil
		})
		close(done)
	}()
	for range 2 {
		select {
		case <-started:
		case <-time.After(time.Second):
			close(gate)
			<-done
			t.Fatal("workers did not run concurrently")
		}
	}
	close(gate)
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("workers did not finish")
	}
	noError(t, err)
	same(t, got, []int{1, 2, 3, 4})
	if peak.Load() > 2 {
		t.Fatalf("peak active = %d", peak.Load())
	}
	equal(t, active.Load(), int64(0))
}

func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
func same(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v; want %#v", got, want)
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
