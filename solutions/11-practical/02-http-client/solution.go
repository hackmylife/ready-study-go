//go:build ignore

package koan

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type Page struct {
	Items []string `json:"items"`
	Next  int      `json:"next"`
}

func List(ctx context.Context, client *http.Client, baseURL string, maxPages int) ([]string, error) {
	if maxPages <= 0 {
		return nil, errors.New("maxPages must be positive")
	}
	base, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	page := 1
	seen := make(map[int]bool)
	items := []string{}
	for count := 0; count < maxPages; count++ {
		if seen[page] {
			return nil, errors.New("pagination cycle")
		}
		seen[page] = true
		query := base.Query()
		query.Set("page", strconv.Itoa(page))
		base.RawQuery = query.Encode()
		req, err := http.NewRequestWithContext(ctx, "GET", base.String(), nil)
		if err != nil {
			return nil, err
		}
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		data, readErr := io.ReadAll(io.LimitReader(resp.Body, 1<<20+1))
		closeErr := resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("status: %d", resp.StatusCode)
		}
		if err := errors.Join(readErr, closeErr); err != nil {
			return nil, err
		}
		if len(data) > 1<<20 {
			return nil, errors.New("response too large")
		}
		var result Page
		if err := json.Unmarshal(data, &result); err != nil {
			return nil, err
		}
		items = append(items, result.Items...)
		if result.Next == 0 {
			return items, nil
		}
		if result.Next < 0 {
			return nil, errors.New("invalid next page")
		}
		page = result.Next
	}
	return nil, errors.New("page limit exceeded")
}
