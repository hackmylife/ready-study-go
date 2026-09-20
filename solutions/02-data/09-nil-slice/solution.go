//go:build ignore

package koan

import (
	"encoding/json"
)

func JSONList(values []string) ([]byte, error) {
	if values == nil {
		values = []string{}
	}
	return json.Marshal(values)
}
