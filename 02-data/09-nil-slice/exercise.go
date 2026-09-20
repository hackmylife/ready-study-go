package koan

import (
	"encoding/json"
)

func JSONList(values []string) ([]byte, error) { // TODO: 常にJSON配列にする
	return json.Marshal(values)
}
