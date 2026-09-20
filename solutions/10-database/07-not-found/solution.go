//go:build ignore

package koan

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

var ErrUserNotFound = errors.New("user not found")

func Find(ctx context.Context, db *sql.DB, id string) (string, error) {
	var name string
	err := db.QueryRowContext(ctx, `SELECT name FROM users WHERE id=$1`, id).Scan(&name)
	if errors.Is(err, sql.ErrNoRows) {
		return "", ErrUserNotFound
	}
	if err != nil {
		return "", fmt.Errorf("find user: %w", err)
	}
	return name, nil
}
