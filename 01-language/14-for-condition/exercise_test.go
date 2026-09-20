package koan

import "testing"

func TestDigitCount(t *testing.T) {
	for _, tt := range []struct{ n, want int }{{0, 1}, {7, 1}, {9, 1}, {10, 2}, {99, 2}, {100, 3}, {1200, 4}} {
		if got := DigitCount(tt.n); got != tt.want {
			t.Errorf("DigitCount(%d) = %d; want %d", tt.n, got, tt.want)
		}
	}
}
