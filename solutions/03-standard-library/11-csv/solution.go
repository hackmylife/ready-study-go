//go:build ignore

package koan

import (
	"encoding/csv"
	"io"
)

func ReadCSV(r io.Reader) ([][]string, error) {
	return csv.NewReader(r).ReadAll()
}
