//go:build ignore

package koan

type MemoryRepository struct{ values map[string]string }

func NewRepository() *MemoryRepository {
	return &MemoryRepository{values: map[string]string{"u1": "Aki"}}
}
func (r *MemoryRepository) Find(id string) (string, bool) { v, ok := r.values[id]; return v, ok }
