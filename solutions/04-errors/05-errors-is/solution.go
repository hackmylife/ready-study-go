//go:build ignore

package koan

import (
	"errors"
	"os"
)

func Missing(err error) bool { return errors.Is(err, os.ErrNotExist) }
