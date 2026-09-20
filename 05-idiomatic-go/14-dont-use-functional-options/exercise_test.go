package koan

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestFormatter(t *testing.T) {
	newFormatter, ok := any(NewFormatter).(func(string) *Formatter)
	if !ok {
		t.Fatal("use one explicit prefix argument")
	}
	equal(t, newFormatter("> ").Format("Go"), "> Go")
	equal(t, newFormatter("").Format("Go"), "Go")
	f, err := parser.ParseFile(token.NewFileSet(), "exercise.go", nil, 0)
	noError(t, err)
	ast.Inspect(f, func(n ast.Node) bool {
		if ts, ok := n.(*ast.TypeSpec); ok && ts.Name.Name == "Option" {
			t.Error("remove Option")
		}
		if fn, ok := n.(*ast.FuncDecl); ok && fn.Name.Name == "WithPrefix" {
			t.Error("remove WithPrefix")
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
