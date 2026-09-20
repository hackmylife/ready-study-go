package koan

import "testing"

func TestSumEvenBeforeNegative(t *testing.T) {
	for _, tt := range []struct {
		values []int
		want   int
	}{
		{nil, 0}, {[]int{}, 0}, {[]int{2, 3, 4, -1, 10}, 6}, {[]int{-2, 8}, 0},
		{[]int{1, 3, 5}, 0}, {[]int{0, 2, 4}, 6}, {[]int{2, -4, 6}, 2},
	} {
		if got := SumEvenBeforeNegative(tt.values); got != tt.want {
			t.Errorf("input=%v: got %d; want %d", tt.values, got, tt.want)
		}
	}
}
