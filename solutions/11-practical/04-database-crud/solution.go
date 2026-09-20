//go:build ignore

package koan

import (
	"context"
	"database/sql"
	"errors"
	"strings"
)

type User struct{ ID, Name string }

func Create(ctx context.Context, db *sql.DB, u User) error {
	if u.ID == "" || strings.TrimSpace(u.Name) == "" {
		return errors.New("id and name are required")
	}
	_, err := db.ExecContext(ctx, `INSERT INTO users(id,name) VALUES($1,$2)`, u.ID, u.Name)
	return err
}
func Find(ctx context.Context, db *sql.DB, id string) (User, error) {
	var u User
	err := db.QueryRowContext(ctx, `SELECT id,name FROM users WHERE id=$1`, id).Scan(&u.ID, &u.Name)
	return u, err
}
func Rename(ctx context.Context, db *sql.DB, id, name string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("name is required")
	}
	result, err := db.ExecContext(ctx, `UPDATE users SET name=$1 WHERE id=$2`, name, id)
	return affected(result, err)
}
func Delete(ctx context.Context, db *sql.DB, id string) error {
	result, err := db.ExecContext(ctx, `DELETE FROM users WHERE id=$1`, id)
	return affected(result, err)
}
func affected(result sql.Result, err error) error {
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return sql.ErrNoRows
	}
	return nil
}
