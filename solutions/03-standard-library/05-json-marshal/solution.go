//go:build ignore

package koan

import (
	"encoding/json"
)

type User struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Secret string `json:"-"`
}

func Encode(u User) ([]byte, error) { return json.Marshal(u) }
