//go:build ignore

package koan

import (
	"context"
	"database/sql"
)

func Rename(ctx context.Context, db *sql.DB, id, name string) error {
	result, err := db.ExecContext(ctx, `UPDATE users SET name=$1 WHERE id=$2`, name, id)
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return sql.ErrNoRows
	}
	return nil
}
