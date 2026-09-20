package koan

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestUpper(t *testing.T) {
	equal(t, Upper("Go"), "GO")
	equal(t, Upper("猫"), "猫")
	f, err := parser.ParseFile(token.NewFileSet(), "exercise.go", nil, 0)
	noError(t, err)
	ast.Inspect(f, func(n ast.Node) bool {
		switch n.(type) {
		case *ast.GoStmt, *ast.ChanType:
			t.Error("remove goroutine and channel")
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
