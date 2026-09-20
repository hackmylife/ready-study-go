package koan

import (
	"context"
	"errors"
	"ready-study-go/internal/dbtest"
	"reflect"
	"testing"
)

func TestNames(t *testing.T) {
	db := dbtest.Open(t)
	db.SetMaxOpenConns(1)
	dbtest.Users(t, db)
	for range 2 {
		names, err := Names(context.Background(), db)
		noError(t, err)
		same(t, names, []string{"Aki", "Ren"})
	}
	dbtest.Exec(t, db, `DELETE FROM users`)
	names, err := Names(context.Background(), db)
	noError(t, err)
	same(t, names, []string{})
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err = Names(ctx, db)
	if !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}
}
func same(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v; want %#v", got, want)
	}
}
func noError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
