package koan

import (
	"context"
	"ready-study-go/internal/dbtest"
	"testing"
)

func TestRenameBoth(t *testing.T) {
	db := dbtest.Open(t)
	dbtest.Users(t, db)
	dbtest.Exec(t, db, `ALTER TABLE users ADD CHECK(name <> '')`)
	noError(t, RenameBoth(context.Background(), db, "Nao", "Mei"))
	var first, second string
	noError(t, db.QueryRow(`SELECT name FROM users WHERE id='u1'`).Scan(&first))
	equal(t, first, "Nao")
	wantError(t, RenameBoth(context.Background(), db, "changed", ""))
	noError(t, db.QueryRow(`SELECT name FROM users WHERE id='u1'`).Scan(&first))
	noError(t, db.QueryRow(`SELECT name FROM users WHERE id='u2'`).Scan(&second))
	equal(t, first, "Nao")
	equal(t, second, "Mei")
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
