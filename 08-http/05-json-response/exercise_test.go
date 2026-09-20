package koan

import (
	"encoding/json"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestUser(t *testing.T) {
	w := httptest.NewRecorder()
	UserHandler(w, httptest.NewRequest("GET", "/", nil))
	equal(t, w.Code, 200)
	equal(t, w.Header().Get("Content-Type"), "application/json")
	var got map[string]string
	noError(t, json.Unmarshal(w.Body.Bytes(), &got))
	same(t, got, map[string]string{"name": "Aki"})
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
