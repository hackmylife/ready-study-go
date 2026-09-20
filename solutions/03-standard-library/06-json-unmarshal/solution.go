//go:build ignore

package koan

import (
	"encoding/json"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Decode(data []byte) (User, error) {
	var u User
	err := json.Unmarshal(data, &u)
	return u, err
}
