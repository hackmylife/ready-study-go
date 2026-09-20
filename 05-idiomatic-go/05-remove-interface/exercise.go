package koan

type Repository interface{ Find(string) (string, bool) }
type MemoryRepository struct{ values map[string]string }

func NewRepository() Repository                           { return &MemoryRepository{values: map[string]string{"u1": "Aki"}} } // TODO: 具体型を返す
func (r *MemoryRepository) Find(id string) (string, bool) { v, ok := r.values[id]; return v, ok }
