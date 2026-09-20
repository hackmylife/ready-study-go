//go:build ignore

package koan

import (
	"os"
)

func Write(path, content string) error {
	return os.WriteFile(path, []byte(content), 0600)
}
