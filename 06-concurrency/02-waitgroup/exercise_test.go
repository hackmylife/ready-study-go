package koan

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestRunAll(t *testing.T) {
	var count atomic.Int64
	release := make(chan struct{})
	started := make(chan struct{}, 3)
	done := make(chan struct{})
	jobs := make([]func(), 3)
	for i := range jobs {
		jobs[i] = func() { started <- struct{}{}; <-release; count.Add(1) }
	}
	go func() { RunAll(jobs); close(done) }()
	defer func() { close(release); <-done }()
	for range jobs {
		select {
		case <-started:
		case <-time.After(time.Second):
			t.Fatal("jobs did not run concurrently")
		}
	}
	select {
	case <-done:
		t.Fatal("returned before jobs finished")
	default:
	}
}
func TestRunAllEmptyAndCount(t *testing.T) {
	RunAll(nil)
	var count atomic.Int64
	RunAll([]func(){func() { count.Add(1) }, func() { count.Add(1) }})
	equal(t, count.Load(), int64(2))
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
