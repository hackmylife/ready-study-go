package koan

import (
	"context"
	"database/sql"
	"errors"
)

var ErrUserNotFound = errors.New("user not found")

func Find(ctx context.Context, db *sql.DB, id string) (string, error) { // TODO: 未登録を用途のerrorにする
	return "", nil
}
