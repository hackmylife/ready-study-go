package koan

import (
	"context"
	"database/sql"
	"errors"
	"ready-study-go/internal/dbtest"
	"reflect"
	"testing"
)

func TestCRUD(t *testing.T) {
	db := dbtest.Open(t)
	dbtest.Users(t, db)
	ctx := context.Background()
	u := User{ID: "u3", Name: "Nao"}
	noError(t, Create(ctx, db, u))
	got, err := Find(ctx, db, "u3")
	noError(t, err)
	same(t, got, u)
	wantError(t, Create(ctx, db, u))
	noError(t, Rename(ctx, db, "u3", "Mei"))
	got, err = Find(ctx, db, "u3")
	noError(t, err)
	equal(t, got.Name, "Mei")
	noError(t, Delete(ctx, db, "u3"))
	_, err = Find(ctx, db, "u3")
	if !errors.Is(err, sql.ErrNoRows) {
		t.Fatal(err)
	}
	if !errors.Is(Delete(ctx, db, "u3"), sql.ErrNoRows) {
		t.Fatal("missing delete")
	}
	if !errors.Is(Rename(ctx, db, "missing", "x"), sql.ErrNoRows) {
		t.Fatal("missing update")
	}
	wantError(t, Create(ctx, db, User{}))
	wantError(t, Rename(ctx, db, "u1", " "))
}

func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
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
func wantError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("expected an error")
	}
}
