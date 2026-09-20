//go:build ignore

package koan

import (
	"bufio"
	"io"
)

func EachLine(r io.Reader, visit func(string) error) error {
	s := bufio.NewScanner(r)
	for s.Scan() {
		if err := visit(s.Text()); err != nil {
			return err
		}
	}
	return s.Err()
}
