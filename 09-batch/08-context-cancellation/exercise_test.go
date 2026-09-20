package koan

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

func TestCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var got []string
	err := Process(ctx, []string{"a", "b"}, func(child context.Context, s string) error {
		equal(t, child, ctx)
		got = append(got, s)
		cancel()
		return nil
	})
	same(t, got, []string{"a"})
	if !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}
	cause := errors.New("failed")
	err = Process(context.Background(), []string{"a"}, func(context.Context, string) error { return cause })
	if !errors.Is(err, cause) {
		t.Fatal(err)
	}
	noError(t, Process(context.Background(), nil, nil))
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
func same(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v; want %#v", got, want)
	}
}
func noError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
