//go:build ignore

package koan

import (
	"bufio"
	"io"
)

type Decoder struct{ scanner *bufio.Scanner }

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{scanner: bufio.NewScanner(r)}
}
func (d *Decoder) Next() (string, bool) {
	ok := d.scanner.Scan()
	return d.scanner.Text(), ok
}
func (d *Decoder) Err() error {
	return d.scanner.Err()
}
