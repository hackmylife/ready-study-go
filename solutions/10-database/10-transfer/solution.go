//go:build ignore

package koan

import (
	"context"
	"database/sql"
	"errors"
	"math"
)

var ErrInvalid = errors.New("invalid transfer")
var ErrFunds = errors.New("insufficient funds")

func Transfer(ctx context.Context, db *sql.DB, from, to string, amount int64) error {
	if from == to || amount <= 0 {
		return ErrInvalid
	}
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	rows, err := tx.QueryContext(ctx, `SELECT id,balance FROM accounts WHERE id IN ($1,$2) ORDER BY id FOR UPDATE`, from, to)
	if err != nil {
		return err
	}
	balances := make(map[string]int64)
	for rows.Next() {
		var id string
		var balance int64
		if err := rows.Scan(&id, &balance); err != nil {
			rows.Close()
			return err
		}
		balances[id] = balance
	}
	err = rows.Err()
	rows.Close()
	if err != nil {
		return err
	}
	if len(balances) != 2 {
		return sql.ErrNoRows
	}
	if balances[from] < amount {
		return ErrFunds
	}
	if balances[to] > math.MaxInt64-amount {
		return ErrInvalid
	}
	if _, err := tx.ExecContext(ctx, `UPDATE accounts SET balance=balance-$1 WHERE id=$2`, amount, from); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, `UPDATE accounts SET balance=balance+$1 WHERE id=$2`, amount, to); err != nil {
		return err
	}
	return tx.Commit()
}
