//go:build ignore

package koan

import (
	"fmt"
)

func Load(path string, read func(string) ([]byte, error)) ([]byte, error) {
	b, err := read(path)
	if err != nil {
		return nil, fmt.Errorf("load %s: %w", path, err)
	}
	return b, nil
}
