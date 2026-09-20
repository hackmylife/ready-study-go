package koan

import (
	"errors"
	"fmt"
)

var ErrMissing = errors.New("missing")

func Load(found bool) error {
	if !found {
		return fmt.Errorf("load: %w", ErrMissing)
	}
	return nil
}
