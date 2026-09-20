package koan

import (
	"context"
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect(ctx context.Context, dsn string) (*sql.DB, error) { // TODO: 接続を開いて確認する
	return nil, nil
}
