package koan

import (
	"io"
)

type CountingWriter struct {
	Writer io.Writer
	Bytes  int
}

func (w *CountingWriter) Write(p []byte) (int, error) { // TODO: 委譲して実際の件数を記録する
	return w.Writer.Write(p)
}
