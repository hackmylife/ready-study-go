package koan

import (
	"os"
)

func Missing(err error) bool { // TODO: 原因をたどって判定する
	return err == os.ErrNotExist
}
