package koan

import (
	"io"
)

type Request struct {
	Name string `json:"name"`
}

func Decode(r io.Reader) (Request, error) { // TODO: 厳密に一つだけ読む
	return Request{}, nil
}
