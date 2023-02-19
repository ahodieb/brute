package slices

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   int
	}{
		{name: "empty", values: []int{}, want: 0},
		{name: "positive values", values: []int{1, 2, 3}, want: 6},
		{name: "negative values", values: []int{-1, -2, -3}, want: -6},
		{name: "mixed values", values: []int{-1, 2, -1}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.values...); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleSum() {
	fmt.Println(Sum(1, 2, 3, 4))
	fmt.Println(Sum(-1, 1))

	// Output:
	// 10
	// 0
}
