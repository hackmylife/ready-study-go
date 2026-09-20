package koan

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestEncode(t *testing.T) {
	b, err := Encode(User{Name: "猫", Age: 0, Secret: "private"})
	noError(t, err)
	var got map[string]any
	noError(t, json.Unmarshal(b, &got))
	same(t, got, map[string]any{"name": "猫", "age": float64(0)})
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
