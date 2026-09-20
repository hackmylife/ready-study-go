package koan

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestMapCancelsAndJoinsWorkers(t *testing.T) {
	cause := errors.New("job failed")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	blocked := make(chan struct{})
	exited := make(chan struct{})
	done := make(chan error, 1)
	go func() {
		_, err := Map(ctx, 2, []int{0, 1, 2}, func(child context.Context, value int) (int, error) {
			if value == 0 {
				close(blocked)
				<-child.Done()
				close(exited)
				return 0, child.Err()
			}
			<-blocked
			return 0, cause
		})
		done <- err
	}()
	select {
	case err := <-done:
		if !errors.Is(err, cause) {
			t.Fatalf("want first failure, got %v", err)
		}
		select {
		case <-exited:
		default:
			t.Fatal("Map returned before the blocked worker exited")
		}
	case <-time.After(time.Second):
		t.Fatal("a failing job did not cancel the blocked worker")
	}
}
