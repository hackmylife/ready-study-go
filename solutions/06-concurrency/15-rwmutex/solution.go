//go:build ignore

package koan

import (
	"sync"
)

type Scores struct {
	mu     sync.RWMutex
	scores map[string]int
}

func (s *Scores) Set(name string, score int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.scores == nil {
		s.scores = make(map[string]int)
	}
	s.scores[name] = score
}

func (s *Scores) Get(name string) (int, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	score, ok := s.scores[name]
	return score, ok
}

func (s *Scores) View(fn func(scores map[string]int)) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	fn(s.scores)
}
