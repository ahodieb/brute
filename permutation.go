package brute

import "golang.org/x/exp/slices"

// Permutations returns all the possible different ordering of specified list of items
func Permutations[T any](items []T) [][]T {
	iterations := permute[T]([]T{}, items)
	return iterations
}

func permute[T any](permutation []T, items []T) [][]T {
	if len(items) == 0 {
		return [][]T{permutation}
	}

	var result [][]T
	for i, item := range items {
		newPermutation := append(permutation, item)
		remainingItems := slices.Delete(slices.Clone(items), i, i+1)
		next := permute(newPermutation, remainingItems)
		result = append(result, next...)
	}

	return result
}
