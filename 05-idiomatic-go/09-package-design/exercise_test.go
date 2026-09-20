package koan

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestInvoice(t *testing.T) {
	equal(t, InvoiceTotal(101, 10), 111)
	equal(t, InvoiceTotal(0, 10), 0)
	f, err := parser.ParseFile(token.NewFileSet(), "exercise.go", nil, 0)
	noError(t, err)
	for _, d := range f.Decls {
		if fn, ok := d.(*ast.FuncDecl); ok && fn.Name.IsExported() && fn.Name.Name != "InvoiceTotal" {
			t.Errorf("internal calculation should not be exported: %s", fn.Name.Name)
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
