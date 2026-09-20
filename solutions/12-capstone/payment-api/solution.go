//go:build ignore

package koan

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"io"
	"log/slog"
	"math"
	"net/http"
	"strconv"
	"strings"
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

func (s *Store) CreateAccount(ctx context.Context, name string, balance int64) (Account, error) {
	name = strings.TrimSpace(name)
	if name == "" || len(name) > 200 || balance < 0 {
		return Account{}, ErrInvalid
	}
	var account Account
	err := s.DB.QueryRowContext(ctx, `INSERT INTO accounts(name,balance) VALUES($1,$2) RETURNING id,name,balance`, name, balance).Scan(&account.ID, &account.Name, &account.Balance)
	return account, err
}

func (s *Store) Account(ctx context.Context, id int64) (Account, error) {
	var account Account
	err := s.DB.QueryRowContext(ctx, `SELECT id,name,balance FROM accounts WHERE id=$1`, id).Scan(&account.ID, &account.Name, &account.Balance)
	if errors.Is(err, sql.ErrNoRows) {
		return Account{}, ErrNotFound
	}
	return account, err
}

func (s *Store) TransferByID(ctx context.Context, id int64) (Transfer, error) {
	var transfer Transfer
	err := s.DB.QueryRowContext(ctx, `SELECT id,from_id,to_id,amount FROM transfers WHERE id=$1`, id).Scan(&transfer.ID, &transfer.FromID, &transfer.ToID, &transfer.Amount)
	if errors.Is(err, sql.ErrNoRows) {
		return Transfer{}, ErrNotFound
	}
	return transfer, err
}

func (s *Store) Send(ctx context.Context, key string, from, to, amount int64) (Transfer, bool, error) {
	if strings.TrimSpace(key) == "" || len(key) > 200 || from <= 0 || to <= 0 || from == to || amount <= 0 {
		return Transfer{}, false, ErrInvalid
	}
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return Transfer{}, false, err
	}
	defer tx.Rollback()
	existing, err := transferByKey(ctx, tx, key)
	if err == nil {
		return matchTransfer(existing, from, to, amount)
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return Transfer{}, false, err
	}

	// Lock accounts by ID, because opposite-direction transfers must acquire locks in the same order.
	rows, err := tx.QueryContext(ctx, `SELECT id,balance FROM accounts WHERE id IN ($1,$2) ORDER BY id FOR UPDATE`, from, to)
	if err != nil {
		return Transfer{}, false, err
	}
	balances := make(map[int64]int64)
	for rows.Next() {
		var id, balance int64
		if err := rows.Scan(&id, &balance); err != nil {
			rows.Close()
			return Transfer{}, false, err
		}
		balances[id] = balance
	}
	err = rows.Err()
	rows.Close()
	if err != nil {
		return Transfer{}, false, err
	}
	if len(balances) != 2 {
		return Transfer{}, false, ErrNotFound
	}

	// Recheck the key after locking, because another transfer may have committed while this one waited.
	existing, err = transferByKey(ctx, tx, key)
	if err == nil {
		return matchTransfer(existing, from, to, amount)
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return Transfer{}, false, err
	}
	if balances[from] < amount {
		return Transfer{}, false, ErrFunds
	}
	if balances[to] > math.MaxInt64-amount {
		return Transfer{}, false, ErrInvalid
	}
	if _, err := tx.ExecContext(ctx, `UPDATE accounts SET balance=balance-$1 WHERE id=$2`, amount, from); err != nil {
		return Transfer{}, false, err
	}
	if _, err := tx.ExecContext(ctx, `UPDATE accounts SET balance=balance+$1 WHERE id=$2`, amount, to); err != nil {
		return Transfer{}, false, err
	}
	var transfer Transfer
	err = tx.QueryRowContext(ctx, `INSERT INTO transfers(idempotency_key,from_id,to_id,amount) VALUES($1,$2,$3,$4) RETURNING id,from_id,to_id,amount`, key, from, to, amount).Scan(&transfer.ID, &transfer.FromID, &transfer.ToID, &transfer.Amount)
	if err != nil {
		var pgerr *pgconn.PgError
		if errors.As(err, &pgerr) && pgerr.Code == "23505" {
			return Transfer{}, false, ErrConflict
		}
		return Transfer{}, false, err
	}
	if err := tx.Commit(); err != nil {
		return Transfer{}, false, err
	}
	return transfer, true, nil
}

func transferByKey(ctx context.Context, tx *sql.Tx, key string) (Transfer, error) {
	var transfer Transfer
	err := tx.QueryRowContext(ctx, `SELECT id,from_id,to_id,amount FROM transfers WHERE idempotency_key=$1`, key).Scan(&transfer.ID, &transfer.FromID, &transfer.ToID, &transfer.Amount)
	return transfer, err
}

func matchTransfer(existing Transfer, from, to, amount int64) (Transfer, bool, error) {
	if existing.FromID != from || existing.ToID != to || existing.Amount != amount {
		return Transfer{}, false, ErrConflict
	}
	return existing, false, nil
}

func Routes(store *Store, logger *slog.Logger) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /accounts", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Name    string `json:"name"`
			Balance int64  `json:"balance"`
		}
		if err := decode(w, r, &req); err != nil {
			writeError(w, r, logger, ErrInvalid)
			return
		}
		account, err := store.CreateAccount(r.Context(), req.Name, req.Balance)
		if err != nil {
			writeError(w, r, logger, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("/accounts/%d", account.ID))
		writeJSON(w, 201, account)
	})
	mux.HandleFunc("GET /accounts/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := parseID(r.PathValue("id"))
		if err != nil {
			writeError(w, r, logger, err)
			return
		}
		account, err := store.Account(r.Context(), id)
		if err != nil {
			writeError(w, r, logger, err)
			return
		}
		writeJSON(w, 200, account)
	})
	mux.HandleFunc("POST /transfers", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			FromID int64 `json:"from_id"`
			ToID   int64 `json:"to_id"`
			Amount int64 `json:"amount"`
		}
		if err := decode(w, r, &req); err != nil {
			writeError(w, r, logger, ErrInvalid)
			return
		}
		transfer, created, err := store.Send(r.Context(), r.Header.Get("Idempotency-Key"), req.FromID, req.ToID, req.Amount)
		if err != nil {
			writeError(w, r, logger, err)
			return
		}
		status := 200
		if created {
			status = 201
		}
		w.Header().Set("Location", fmt.Sprintf("/transfers/%d", transfer.ID))
		writeJSON(w, status, transfer)
	})
	mux.HandleFunc("GET /transfers/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := parseID(r.PathValue("id"))
		if err != nil {
			writeError(w, r, logger, err)
			return
		}
		transfer, err := store.TransferByID(r.Context(), id)
		if err != nil {
			writeError(w, r, logger, err)
			return
		}
		writeJSON(w, 200, transfer)
	})
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.InfoContext(r.Context(), "request", "method", r.Method, "path", r.URL.Path)
		mux.ServeHTTP(w, r)
	})
}

func parseID(raw string) (int64, error) {
	id, err := strconv.ParseInt(raw, 10, 64)
	if err != nil || id <= 0 {
		return 0, ErrInvalid
	}
	return id, nil
}
func decode(w http.ResponseWriter, r *http.Request, target any) error {
	r.Body = http.MaxBytesReader(w, r.Body, 4096)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(target); err != nil {
		return err
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		return ErrInvalid
	}
	return nil
}
func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}
func writeError(w http.ResponseWriter, r *http.Request, logger *slog.Logger, err error) {
	status, code := 500, "internal_error"
	switch {
	case errors.Is(err, ErrInvalid):
		status, code = 400, "invalid_request"
	case errors.Is(err, ErrNotFound):
		status, code = 404, "not_found"
	case errors.Is(err, ErrFunds):
		status, code = 409, "insufficient_funds"
	case errors.Is(err, ErrConflict):
		status, code = 409, "idempotency_conflict"
	}
	if status == 500 {
		logger.ErrorContext(r.Context(), "request failed", "error", err)
	}
	writeJSON(w, status, struct {
		Error string `json:"error"`
	}{Error: code})
}
