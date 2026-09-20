//go:build ignore

package koan

import (
	"io"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) { _, _ = io.WriteString(w, "Hello, Go\n") }
