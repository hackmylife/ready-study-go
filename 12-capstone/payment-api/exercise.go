package koan

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"net/http"
)

type Account struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Balance int64  `json:"balance"`
}
type Transfer struct {
	ID     int64 `json:"id"`
	FromID int64 `json:"from_id"`
	ToID   int64 `json:"to_id"`
	Amount int64 `json:"amount"`
}
type Store struct{ DB *sql.DB }

var ErrInvalid = errors.New("invalid request")
var ErrNotFound = errors.New("not found")
var ErrFunds = errors.New("insufficient funds")
var ErrConflict = errors.New("idempotency key conflict")

func (s *Store) CreateAccount(ctx context.Context, name string, balance int64) (Account, error) { // TODO: 口座を検証して作る
	return Account{}, nil
}
func (s *Store) Account(ctx context.Context, id int64) (Account, error) { // TODO: 口座を取得する
	return Account{}, nil
}
func (s *Store) TransferByID(ctx context.Context, id int64) (Transfer, error) { // TODO: 送金を取得する
	return Transfer{}, nil
}
func (s *Store) Send(ctx context.Context, key string, from, to, amount int64) (Transfer, bool, error) { // TODO: 原子的かつ冪等に送金する
	return Transfer{}, false, nil
}
func Routes(store *Store, logger *slog.Logger) http.Handler { // TODO: JSON APIのルートとエラー応答を実装する
	return http.NewServeMux()
}
