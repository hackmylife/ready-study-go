package koan

import (
	"context"
	"database/sql"
)

func WithTx(ctx context.Context, db *sql.DB, fn func(*sql.Tx) error) error { // TODO: commitとrollbackを管理する
	return nil
}
