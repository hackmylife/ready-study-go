//go:build ignore

package koan

import (
	"errors"
	"io"
)

func ReadAndClose(r io.ReadCloser) (data []byte, err error) {
	defer func() { err = errors.Join(err, r.Close()) }()
	return io.ReadAll(r)
}
