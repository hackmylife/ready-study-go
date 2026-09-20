//go:build ignore

package koan

type Set struct{ values map[string]struct{} }

func (s *Set) Add(value string) {
	if s.values == nil {
		s.values = make(map[string]struct{})
	}
	s.values[value] = struct{}{}
}
func (s *Set) Has(value string) bool { _, ok := s.values[value]; return ok }
