package koan

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestCanShip(t *testing.T) {
	noError(t, CanShip(true, 1))
	for _, tt := range []struct {
		paid  bool
		stock int
		want  string
	}{{false, 1, "unpaid"}, {false, 0, "unpaid"}, {true, 0, "out of stock"}, {true, -1, "out of stock"}} {
		err := CanShip(tt.paid, tt.stock)
		wantError(t, err)
		equal(t, err.Error(), tt.want)
	}
	f, err := parser.ParseFile(token.NewFileSet(), "exercise.go", nil, 0)
	noError(t, err)
	ast.Inspect(f, func(n ast.Node) bool {
		if stmt, ok := n.(*ast.IfStmt); ok && stmt.Else != nil {
			t.Error("replace else branches with early returns")
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
func wantError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("expected an error")
	}
}
