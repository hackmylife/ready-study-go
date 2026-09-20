package koan

import (
	"context"
	"database/sql"
)

func Email(ctx context.Context, db *sql.DB, id string) (string, bool, error) { // TODO: NULLを扱う
	return "", false, nil
}
