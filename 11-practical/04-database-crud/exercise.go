package koan

import (
	"context"
	"database/sql"
)

type User struct{ ID, Name string }

func Create(ctx context.Context, db *sql.DB, u User) error { // TODO: ユーザーを作成する
	return nil
}
func Find(ctx context.Context, db *sql.DB, id string) (User, error) { return User{}, nil }
func Rename(ctx context.Context, db *sql.DB, id, name string) error { return nil }
func Delete(ctx context.Context, db *sql.DB, id string) error       { return nil }
