//go:build ignore

package koan

import (
	"encoding/json"
	"errors"
	"io"
)

type Request struct {
	Name string `json:"name"`
}

func Decode(r io.Reader) (Request, error) {
	var req Request
	d := json.NewDecoder(r)
	d.DisallowUnknownFields()
	if err := d.Decode(&req); err != nil {
		return Request{}, err
	}
	var extra any
	if err := d.Decode(&extra); err != io.EOF {
		if err == nil {
			return Request{}, errors.New("multiple JSON values")
		}
		return Request{}, err
	}
	return req, nil
}
