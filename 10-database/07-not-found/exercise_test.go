package koan

import (
	"context"
	"errors"
	"ready-study-go/internal/dbtest"
	"testing"
)

func TestFind(t *testing.T) {
	db := dbtest.Open(t)
	dbtest.Users(t, db)
	got, err := Find(context.Background(), db, "u1")
	noError(t, err)
	equal(t, got, "Aki")
	_, err = Find(context.Background(), db, "missing")
	if !errors.Is(err, ErrUserNotFound) {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err = Find(ctx, db, "u1")
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("DB failure is not not-found: %v", err)
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
