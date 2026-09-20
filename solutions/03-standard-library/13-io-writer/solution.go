//go:build ignore

package koan

import (
	"fmt"
	"io"
)

func WriteGreeting(w io.Writer, name string) error {
	_, err := fmt.Fprintf(w, "Hello, %s\n", name)
	return err
}
