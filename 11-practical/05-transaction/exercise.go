package koan

import (
	"context"
	"database/sql"
	"errors"
)

var ErrInvalid = errors.New("invalid transfer")
var ErrFunds = errors.New("insufficient funds")

func Transfer(ctx context.Context, db *sql.DB, from, to string, amount int64) error { // TODO: 残高を原子的に移動する
	return nil
}
