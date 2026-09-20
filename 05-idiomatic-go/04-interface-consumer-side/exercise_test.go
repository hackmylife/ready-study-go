package koan

import (
	"context"
	"errors"
	"testing"
)

type directory struct {
	ctx context.Context
	id  string
	err error
}

func (d *directory) Name(ctx context.Context, id string) (string, error) {
	d.ctx = ctx
	d.id = id
	return "Aki", d.err
}
func TestGreeting(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	d := &directory{}
	got, err := Greeting(ctx, d, "u1")
	noError(t, err)
	equal(t, got, "Hello, Aki")
	equal(t, d.ctx, ctx)
	equal(t, d.id, "u1")
	d.err = errors.New("offline")
	_, err = Greeting(ctx, d, "u2")
	if !errors.Is(err, d.err) {
		t.Fatal("error lost")
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
