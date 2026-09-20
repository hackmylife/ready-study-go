package koan

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	var c Counter
	var wg sync.WaitGroup
	for range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 100 {
				c.Add(1)
				_ = c.Value()
			}
		}()
	}
	wg.Wait()
	equal(t, c.Value(), 10000)
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
