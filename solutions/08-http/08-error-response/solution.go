//go:build ignore

package koan

import (
	"encoding/json"
	"errors"
	"net/http"
)

var ErrNotFound = errors.New("not found")

func WriteError(w http.ResponseWriter, err error) {
	status, code := 500, "internal_error"
	if errors.Is(err, ErrNotFound) {
		status, code = 404, "not_found"
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(struct {
		Error string `json:"error"`
	}{Error: code})
}
