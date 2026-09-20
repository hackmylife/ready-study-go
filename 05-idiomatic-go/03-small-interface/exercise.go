package koan

import (
	"io"
)

func Send(w io.ReadWriteCloser, message string) error { // TODO: 必要なinterfaceへ絞る
	_, err := io.WriteString(w, message)
	return err
}
