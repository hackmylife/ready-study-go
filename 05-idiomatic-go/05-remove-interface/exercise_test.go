package koan

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestRepository(t *testing.T) {
	if _, ok := any(NewRepository).(func() *MemoryRepository); !ok {
		t.Fatal("return the concrete type")
	}
	r := NewRepository()
	name, ok := r.Find("u1")
	equal(t, name, "Aki")
	equal(t, ok, true)
	_, ok = r.Find("missing")
	equal(t, ok, false)
	f, err := parser.ParseFile(token.NewFileSet(), "exercise.go", nil, 0)
	noError(t, err)
	ast.Inspect(f, func(n ast.Node) bool {
		if ts, ok := n.(*ast.TypeSpec); ok && ts.Name.Name == "Repository" {
			t.Error("remove the unused Repository abstraction")
		}
		return true
	})
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
func noError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
