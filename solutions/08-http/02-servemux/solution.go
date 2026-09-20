//go:build ignore

package koan

import (
	"io"
	"net/http"
)

func Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) { _, _ = io.WriteString(w, "ok") })
	return mux
}
