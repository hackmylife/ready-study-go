package koan

type Set struct{ values map[string]struct{} }

func (s *Set) Add(value string) { // TODO: ゼロ値から追加できるようにする
}
func (s *Set) Has(value string) bool { return false }
