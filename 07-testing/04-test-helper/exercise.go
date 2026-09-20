package koan

import (
	"strconv"
)

func ParseBool(s string) (bool, error) { return strconv.ParseBool(s) }
