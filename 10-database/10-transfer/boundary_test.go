package koan

import (
	"context"
	"errors"
	"math"
	"testing"
	"time"

	"ready-study-go/internal/dbtest"
)

func TestTransferOverflow(t *testing.T) {
	db := dbtest.Open(t)
	dbtest.Accounts(t, db)
	dbtest.Exec(t, db, `UPDATE accounts SET balance=$1 WHERE id=$2`, int64(math.MaxInt64), "b")
	err := Transfer(context.Background(), db, "a", "b", 1)
	if !errors.Is(err, ErrInvalid) {
		t.Fatalf("overflow must be invalid, got %v", err)
	}
	equal(t, dbtest.Balance(t, db, "a"), int64(100))
	equal(t, dbtest.Balance(t, db, "b"), int64(math.MaxInt64))
}

func TestOppositeDirectionTransfers(t *testing.T) {
	db := dbtest.Open(t)
	dbtest.Accounts(t, db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	errs := make(chan error, 2)
	go func() { errs <- Transfer(ctx, db, "a", "b", 10) }()
	go func() { errs <- Transfer(ctx, db, "b", "a", 20) }()
	for range 2 {
		noError(t, <-errs)
	}
	equal(t, dbtest.Balance(t, db, "a"), int64(110))
	equal(t, dbtest.Balance(t, db, "b"), int64(40))
}
