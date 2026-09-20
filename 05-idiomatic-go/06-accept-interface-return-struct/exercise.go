package koan

import (
	"bufio"
	"io"
)

type Decoder struct{ scanner *bufio.Scanner }

func NewDecoder(r io.Reader) *Decoder { // TODO: 具体型を初期化する
	return &Decoder{}
}
func (d *Decoder) Next() (string, bool) {
	if d.scanner == nil {
		return "", false
	}
	ok := d.scanner.Scan()
	return d.scanner.Text(), ok
}
func (d *Decoder) Err() error {
	if d.scanner == nil {
		return nil
	}
	return d.scanner.Err()
}
