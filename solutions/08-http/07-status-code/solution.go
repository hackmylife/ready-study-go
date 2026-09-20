//go:build ignore

package koan

import (
	"net/http"
)

func Created(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "/users/u1")
	w.WriteHeader(http.StatusCreated)
}
