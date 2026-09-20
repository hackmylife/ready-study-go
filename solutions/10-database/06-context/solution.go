//go:build ignore

package koan

import (
	"context"
	"database/sql"
)

func Ping(ctx context.Context, db *sql.DB) error {
	return db.PingContext(ctx)
}
