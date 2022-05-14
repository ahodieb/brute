package brute

import (
	"reflect"
	"testing"
)

func TestPermutations(t *testing.T) {
	type test[T any] struct {
		name string
		args []T
		want [][]T
	}

	tests := []test[int]{
		{
			name: "empty slice",
			args: []int{},
			want: [][]int{{}},
		},
		{
			name: "single item",
			args: []int{1},
			want: [][]int{{1}},
		},
		{
			name: "two items",
			args: []int{1, 2},
			want: [][]int{{1, 2}, {2, 1}},
		},
		{
			name: "three items",
			args: []int{1, 2, 3},
			want: [][]int{
				{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Permutations(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Permutations() = %v, want %v", got, tt.want)
			}
		})
	}
}
