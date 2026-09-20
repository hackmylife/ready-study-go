//go:build ignore

package koan

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"math"
	"slices"
	"strconv"
	"strings"
)

func Run(r io.Reader, w io.Writer, prices map[string]int64) (err error) {
	reader := csv.NewReader(r)
	header, err := reader.Read()
	if err != nil {
		return err
	}
	if !slices.Equal(header, []string{"name", "quantity"}) {
		return errors.New("invalid header")
	}
	writer := csv.NewWriter(w)
	defer func() { writer.Flush(); err = errors.Join(err, writer.Error()) }()
	if err := writer.Write([]string{"name", "quantity", "total"}); err != nil {
		return err
	}
	line := 1
	for {
		row, err := reader.Read()
		if err == io.EOF {
			return nil
		}
		line++
		if err != nil {
			return fmt.Errorf("line %d: %w", line, err)
		}
		name := strings.ToLower(strings.TrimSpace(row[0]))
		quantity, err := strconv.ParseInt(row[1], 10, 64)
		if err != nil {
			return fmt.Errorf("line %d: %w", line, err)
		}
		price, ok := prices[name]
		if !ok || price < 0 || quantity <= 0 || price > math.MaxInt64/quantity {
			return fmt.Errorf("line %d: invalid order", line)
		}
		if err := writer.Write([]string{name, strconv.FormatInt(quantity, 10), strconv.FormatInt(price*quantity, 10)}); err != nil {
			return err
		}
	}
}
