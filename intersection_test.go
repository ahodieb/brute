package brute

import (
	"reflect"
	"testing"
)

func TestIntersections(t *testing.T) {
	tests := []struct {
		name string
		args [][]int
		want []int
	}{
		{
			name: "empty slice",
			args: [][]int{{}},
			want: []int{},
		},
		{
			name: "single slice",
			args: [][]int{{1, 2, 3}},
			want: []int{1, 2, 3},
		},
		{
			name: "two slices all items are common",
			args: [][]int{{1, 2, 3}, {1, 2, 3}},
			want: []int{1, 2, 3},
		},
		{
			name: "two slices with 1 common item",
			args: [][]int{{1, 2, 3}, {3, 4, 5}},
			want: []int{3},
		},
		{
			name: "two slices with 1 common item repeated",
			args: [][]int{{1, 2, 3}, {3, 3, 3}},
			want: []int{3},
		},
		{
			name: "two slices with no common item",
			args: [][]int{{1, 2, 3}, {5, 6, 7}},
			want: []int{},
		},
		{
			name: "three slices with some common item",
			args: [][]int{{1, 2, 3}, {3, 4, 5}, {3, 4, 5}},
			want: []int{3},
		},
		{
			name: "three slices with no common item",
			args: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersection(tt.args...); !reflect.DeepEqual(got, tt.want) && len(got) == 0 && len(tt.want) != 0 {
				t.Errorf("Intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
