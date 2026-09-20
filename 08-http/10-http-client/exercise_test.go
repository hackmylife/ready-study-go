package koan

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
)

type roundTrip func(*http.Request) (*http.Response, error)

func (f roundTrip) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

type trackedBody struct {
	io.Reader
	closed bool
}

func (b *trackedBody) Close() error { b.closed = true; return nil }
func TestFetch(t *testing.T) {
	for _, status := range []int{200, 503} {
		body := &trackedBody{Reader: strings.NewReader("Go")}
		ctx := context.Background()
		client := &http.Client{Transport: roundTrip(func(r *http.Request) (*http.Response, error) {
			equal(t, r.Method, "GET")
			equal(t, r.Context(), ctx)
			return &http.Response{StatusCode: status, Body: body, Header: make(http.Header)}, nil
		})}
		got, err := Fetch(ctx, client, "http://example.test/data")
		if status == 200 {
			noError(t, err)
			equal(t, got, "Go")
		} else {
			wantError(t, err)
		}
		equal(t, body.closed, true)
	}
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
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
