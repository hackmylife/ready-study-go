package koan

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestGreeting(t *testing.T) {
	equal(t, Greeting("Aki"), "Hello, Aki")
	equal(t, LoudGreeting("Aki"), "HELLO, AKI")
	f, err := parser.ParseFile(token.NewFileSet(), "exercise.go", nil, 0)
	noError(t, err)
	for _, d := range f.Decls {
		if fn, ok := d.(*ast.FuncDecl); ok && fn.Name.Name == "Format" {
			t.Fatal("remove the boolean-mode API")
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
