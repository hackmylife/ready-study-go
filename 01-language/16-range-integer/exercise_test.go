package koan

import (
	"reflect"
	"testing"
)

func TestIndices(t *testing.T) {
	for _, tt := range []struct {
		n    int
		want []int
	}{{-2, nil}, {0, nil}, {1, []int{0}}, {4, []int{0, 1, 2, 3}}} {
		if got := Indices(tt.n); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Indices(%d) = %v; want %v", tt.n, got, tt.want)
		}
	}
}
