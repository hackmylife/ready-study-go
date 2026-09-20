package koan

import (
	"io"
)

func ReadAndClose(r io.ReadCloser) ([]byte, error) { // TODO: 読み取りとcloseを扱う
	return io.ReadAll(r)
}
