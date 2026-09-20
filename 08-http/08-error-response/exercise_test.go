package koan

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWriteError(t *testing.T) {
	for _, tt := range []struct {
		err    error
		status int
		code   string
	}{{fmt.Errorf("load: %w", ErrNotFound), 404, "not_found"}, {errors.New("password=secret"), 500, "internal_error"}} {
		w := httptest.NewRecorder()
		WriteError(w, tt.err)
		equal(t, w.Code, tt.status)
		equal(t, w.Header().Get("Content-Type"), "application/json")
		var got map[string]string
		noError(t, json.Unmarshal(w.Body.Bytes(), &got))
		equal(t, got["error"], tt.code)
		if strings.Contains(w.Body.String(), "secret") {
			t.Fatal("internal error leaked")
		}
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
