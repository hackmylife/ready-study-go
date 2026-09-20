package koan

import (
	"context"
	"database/sql"
)

func Name(ctx context.Context, db *sql.DB, id string) (string, error) { // TODO: 単一行を読む
	return "", nil
}
