//go:build ignore

package koan

import (
	"io"
)

type CountingWriter struct {
	Writer io.Writer
	Bytes  int
}

func (w *CountingWriter) Write(p []byte) (int, error) {
	n, err := w.Writer.Write(p)
	w.Bytes += n
	return n, err
}
