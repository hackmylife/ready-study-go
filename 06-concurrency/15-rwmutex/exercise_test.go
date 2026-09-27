package koan

import (
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestScores(t *testing.T) {
	var s Scores
	if _, ok := s.Get("go"); ok {
		t.Fatal("empty Scores has go")
	}
	s.Set("go", 90)
	s.Set("go", 95)
	got, ok := s.Get("go")
	equal(t, got, 95)
	equal(t, ok, true)
}

func TestViewAllowsConcurrentReaders(t *testing.T) {
	var s Scores
	s.Set("go", 90)
	holding := make(chan struct{})
	release := make(chan struct{})
	defer close(release)
	go s.View(func(map[string]int) {
		close(holding)
		<-release
	})
	select {
	case <-holding:
	case <-time.After(time.Second):
		t.Fatal("View did not call fn")
	}
	second := make(chan int)
	go s.View(func(scores map[string]int) { second <- scores["go"] })
	select {
	case got := <-second:
		equal(t, got, 90)
	case <-time.After(time.Second):
		t.Fatal("second View waited while the first View was reading")
	}
}

func TestScoresConcurrentAccess(t *testing.T) {
	var s Scores
	var wg sync.WaitGroup
	for i := range 100 {
		wg.Go(func() {
			s.Set(strconv.Itoa(i%10), i)
			s.Get("0")
			s.View(func(map[string]int) {})
		})
	}
	wg.Wait()
	s.View(func(scores map[string]int) { equal(t, len(scores), 10) })
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
