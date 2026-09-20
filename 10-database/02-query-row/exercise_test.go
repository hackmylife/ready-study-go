package koan

import (
	"context"
	"database/sql"
	"errors"
	"ready-study-go/internal/dbtest"
	"testing"
)

func TestName(t *testing.T) {
	db := dbtest.Open(t)
	dbtest.Users(t, db)
	name, err := Name(context.Background(), db, "u1")
	noError(t, err)
	equal(t, name, "Aki")
	_, err = Name(context.Background(), db, "missing")
	if !errors.Is(err, sql.ErrNoRows) {
		t.Fatal(err)
	}
	_, err = Name(context.Background(), db, "u1' OR 1=1 --")
	if !errors.Is(err, sql.ErrNoRows) {
		t.Fatal("use SQL parameters")
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
