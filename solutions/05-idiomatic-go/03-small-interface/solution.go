//go:build ignore

package koan

import (
	"io"
)

func Send(w io.Writer, message string) error { _, err := io.WriteString(w, message); return err }
