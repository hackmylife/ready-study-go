package koan

import (
	"context"
	"database/sql"
	"errors"
	"ready-study-go/internal/dbtest"
	"testing"
)

func TestRename(t *testing.T) {
	db := dbtest.Open(t)
	dbtest.Users(t, db)
	noError(t, Rename(context.Background(), db, "u1", "Nao"))
	var got string
	noError(t, db.QueryRow(`SELECT name FROM users WHERE id=$1`, "u1").Scan(&got))
	equal(t, got, "Nao")
	if !errors.Is(Rename(context.Background(), db, "missing", "X"), sql.ErrNoRows) {
		t.Fatal("missing row was not reported")
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
