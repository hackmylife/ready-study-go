package koan

import (
	"context"
	"database/sql"
	"errors"
	"ready-study-go/internal/dbtest"
	"testing"
)

func TestEmail(t *testing.T) {
	db := dbtest.Open(t)
	dbtest.Exec(t, db, `CREATE TABLE contacts(id text PRIMARY KEY,email text)`)
	dbtest.Exec(t, db, `INSERT INTO contacts VALUES($1,$2),($3,$4),($5,$6)`, "null", nil, "empty", "", "set", "a@example.test")
	for _, tt := range []struct {
		id, want string
		valid    bool
	}{{"null", "", false}, {"empty", "", true}, {"set", "a@example.test", true}} {
		s, ok, err := Email(context.Background(), db, tt.id)
		noError(t, err)
		equal(t, s, tt.want)
		equal(t, ok, tt.valid)
	}
	_, _, err := Email(context.Background(), db, "missing")
	if !errors.Is(err, sql.ErrNoRows) {
		t.Fatal(err)
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
