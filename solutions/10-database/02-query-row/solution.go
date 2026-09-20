//go:build ignore

package koan

import (
	"context"
	"database/sql"
)

func Name(ctx context.Context, db *sql.DB, id string) (string, error) {
	var name string
	err := db.QueryRowContext(ctx, `SELECT name FROM users WHERE id=$1`, id).Scan(&name)
	return name, err
}
