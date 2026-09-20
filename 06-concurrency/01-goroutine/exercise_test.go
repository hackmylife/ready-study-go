package koan

import (
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	entered := make(chan struct{})
	release := make(chan struct{})
	defer close(release)
	done := Start(func() { close(entered); <-release })
	select {
	case <-entered:
	case <-time.After(time.Second):
		t.Fatal("work did not start")
	}
	select {
	case <-done:
		t.Fatal("finished before work returned")
	default:
	}
}
func TestStartCompletes(t *testing.T) {
	done := Start(func() {})
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("completion was not signaled")
	}
}
