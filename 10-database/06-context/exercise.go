package koan

import (
	"context"
	"database/sql"
)

func Ping(ctx context.Context, db *sql.DB) error { // TODO: context付き操作を使う
	return db.Ping()
}
