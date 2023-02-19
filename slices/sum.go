package slices

import "golang.org/x/exp/constraints"

func Sum[T constraints.Integer](values ...T) T {
	var s T
	for _, v := range values {
		s += v
	}

	return s
}
