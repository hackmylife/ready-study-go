// Package dbtest provides isolated PostgreSQL fixtures for the database exercises.
package dbtest

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"net/url"
	"os"
	"testing"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Fixture struct {
	DB  *sql.DB
	DSN string
}

func Open(t *testing.T) *sql.DB {
	t.Helper()
	return New(t).DB
}

func New(t *testing.T) Fixture {
	t.Helper()
	dsn := os.Getenv("KOANS_DATABASE_URL")
	if dsn == "" {
		t.Skip("PostgreSQL未接続: docs/database.mdを参照しKOANS_DATABASE_URLを設定してください")
	}
	u, err := url.Parse(dsn)
	if err != nil || (u.Scheme != "postgres" && u.Scheme != "postgresql") {
		t.Fatal("KOANS_DATABASE_URL must be a postgres:// or postgresql:// URL")
	}
	admin, err := sql.Open("pgx", dsn)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { admin.Close() })
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var random [12]byte
	if _, err := rand.Read(random[:]); err != nil {
		t.Fatal(err)
	}
	schema := "koans_" + hex.EncodeToString(random[:])
	if _, err := admin.ExecContext(ctx, `CREATE SCHEMA "`+schema+`"`); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if _, err := admin.ExecContext(ctx, `DROP SCHEMA "`+schema+`" CASCADE`); err != nil {
			t.Errorf("cleanup %s: %v", schema, err)
		}
	})
	query := u.Query()
	query.Set("search_path", schema)
	u.RawQuery = query.Encode()
	db, err := sql.Open("pgx", u.String())
	if err != nil {
		t.Fatal(err)
	}
	db.SetMaxOpenConns(4)
	t.Cleanup(func() { db.Close() })
	if err := db.PingContext(ctx); err != nil {
		t.Fatal(err)
	}
	return Fixture{DB: db, DSN: u.String()}
}

func Exec(t *testing.T, db *sql.DB, query string, args ...any) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if _, err := db.ExecContext(ctx, query, args...); err != nil {
		t.Fatal(err)
	}
}

func Users(t *testing.T, db *sql.DB) {
	t.Helper()
	Exec(t, db, `CREATE TABLE users (id text PRIMARY KEY, name text NOT NULL)`)
	Exec(t, db, `INSERT INTO users (id,name) VALUES ($1,$2),($3,$4)`, "u1", "Aki", "u2", "Ren")
}

func Accounts(t *testing.T, db *sql.DB) {
	t.Helper()
	Exec(t, db, `CREATE TABLE accounts (id text PRIMARY KEY, balance bigint NOT NULL CHECK (balance >= 0))`)
	Exec(t, db, `INSERT INTO accounts (id,balance) VALUES ($1,$2),($3,$4)`, "a", 100, "b", 50)
}

func Balance(t *testing.T, db *sql.DB, id string) int64 {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var balance int64
	if err := db.QueryRowContext(ctx, `SELECT balance FROM accounts WHERE id=$1`, id).Scan(&balance); err != nil {
		t.Fatal(fmt.Errorf("balance %s: %w", id, err))
	}
	return balance
}
