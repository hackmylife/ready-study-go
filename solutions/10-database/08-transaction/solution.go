//go:build ignore

package koan

import (
	"context"
	"database/sql"
)

func RenameBoth(ctx context.Context, db *sql.DB, first, second string) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if _, err := tx.ExecContext(ctx, `UPDATE users SET name=$1 WHERE id=$2`, first, "u1"); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, `UPDATE users SET name=$1 WHERE id=$2`, second, "u2"); err != nil {
		return err
	}
	return tx.Commit()
}
