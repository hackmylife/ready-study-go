//go:build ignore

package koan

import (
	"fmt"
)

type Status int

const (
	StatusPending Status = iota
	StatusActive
	StatusClosed
)

func (s Status) String() string {
	switch s {
	case StatusPending:
		return "pending"
	case StatusActive:
		return "active"
	case StatusClosed:
		return "closed"
	}
	return fmt.Sprintf("Status(%d)", int(s))
}
