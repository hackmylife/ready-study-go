package koan

import (
	"context"
	"database/sql"
)

func Rename(ctx context.Context, db *sql.DB, id, name string) error { // TODO: 更新件数を確認する
	return nil
}
