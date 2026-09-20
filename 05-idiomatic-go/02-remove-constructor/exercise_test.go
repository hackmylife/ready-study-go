package koan

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestCounter(t *testing.T) {
	var c Counter
	c.Add(2)
	equal(t, c.Value(), 2)
	f, err := parser.ParseFile(token.NewFileSet(), "exercise.go", nil, 0)
	noError(t, err)
	for _, d := range f.Decls {
		if fn, ok := d.(*ast.FuncDecl); ok && fn.Name.Name == "NewCounter" {
			t.Fatal("remove NewCounter: Counter already has a useful zero value")
		}
	}
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
