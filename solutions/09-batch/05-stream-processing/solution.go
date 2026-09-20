//go:build ignore

package koan

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Transform(r io.Reader, w io.Writer) error {
	s := bufio.NewScanner(r)
	for s.Scan() {
		if _, err := fmt.Fprintln(w, strings.ToUpper(s.Text())); err != nil {
			return err
		}
	}
	return s.Err()
}
