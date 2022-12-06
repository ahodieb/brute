package brute

// Intersection finds the common values found in all items
func Intersection[T comparable](slices ...[]T) []T {
	if len(slices) == 0 {
		return nil
	}

	hash := make(map[T]struct{})
	for i := 0; i < len(slices[0]); i++ {
		hash[slices[0][i]] = struct{}{}
	}

	for _, slice := range slices {
		intersection := make(map[T]struct{})

		for i := 0; i < len(slice); i++ {
			if _, ok := hash[slice[i]]; ok {
				intersection[slice[i]] = struct{}{}
			}
		}
		hash = intersection
	}

	var intersection []T
	for k := range hash {
		intersection = append(intersection, k)
	}

	return intersection
}
