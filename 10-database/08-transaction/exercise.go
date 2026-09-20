package koan

import (
	"context"
	"database/sql"
)

func RenameBoth(ctx context.Context, db *sql.DB, first, second string) error { // TODO: 二更新をtransactionにする
	return nil
}
