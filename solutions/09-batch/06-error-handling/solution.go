//go:build ignore

package koan

import (
	"bufio"
	"fmt"
	"io"
)

func Process(r io.Reader, handle func(string) error) error {
	s := bufio.NewScanner(r)
	line := 0
	for s.Scan() {
		line++
		if err := handle(s.Text()); err != nil {
			return fmt.Errorf("line %d: %w", line, err)
		}
	}
	return s.Err()
}
