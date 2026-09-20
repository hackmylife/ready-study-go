package koan

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestList(t *testing.T) {
	var pages []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pages = append(pages, r.URL.Query().Get("page"))
		equal(t, r.URL.Query().Get("filter"), "go")
		if r.URL.Query().Get("page") == "1" {
			_, _ = io.WriteString(w, `{"items":["a","b"],"next":2}`)
		} else {
			_, _ = io.WriteString(w, `{"items":["c"],"next":0}`)
		}
	}))
	defer server.Close()
	got, err := List(context.Background(), server.Client(), server.URL+"?filter=go", 3)
	noError(t, err)
	same(t, got, []string{"a", "b", "c"})
	same(t, pages, []string{"1", "2"})
}
func TestListFailures(t *testing.T) {
	for _, body := range []string{`{"items":[],"next":1}`, `{"items":[],"next":-1}`, `not json`} {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { _, _ = io.WriteString(w, body) }))
		_, err := List(context.Background(), server.Client(), server.URL, 2)
		server.Close()
		wantError(t, err)
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(503) }))
	defer server.Close()
	_, err := List(context.Background(), server.Client(), server.URL, 1)
	wantError(t, err)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err = List(ctx, server.Client(), server.URL, 1)
	if !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}
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
