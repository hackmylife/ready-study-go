package koan

import (
	"context"
	"database/sql"
	"errors"
	"ready-study-go/internal/dbtest"
	"testing"
)

func TestWithTx(t *testing.T) {
	db := dbtest.Open(t)
	dbtest.Accounts(t, db)
	cause := errors.New("abort")
	err := WithTx(context.Background(), db, func(tx *sql.Tx) error {
		_, err := tx.Exec(`UPDATE accounts SET balance=0 WHERE id='a'`)
		if err != nil {
			return err
		}
		return cause
	})
	if !errors.Is(err, cause) {
		t.Fatal(err)
	}
	equal(t, dbtest.Balance(t, db, "a"), int64(100))
	noError(t, WithTx(context.Background(), db, func(tx *sql.Tx) error { _, err := tx.Exec(`UPDATE accounts SET balance=80 WHERE id='a'`); return err }))
	equal(t, dbtest.Balance(t, db, "a"), int64(80))
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
