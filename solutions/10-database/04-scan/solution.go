//go:build ignore

package koan

import (
	"context"
	"database/sql"
)

func Email(ctx context.Context, db *sql.DB, id string) (string, bool, error) {
	var email sql.NullString
	if err := db.QueryRowContext(ctx, `SELECT email FROM contacts WHERE id=$1`, id).Scan(&email); err != nil {
		return "", false, err
	}
	return email.String, email.Valid, nil
}
