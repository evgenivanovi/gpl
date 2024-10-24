package slice

// Intersection returns intersection for slices of various built-in types.
func Intersection[E comparable](one []E, two []E) []E {
	if len(one) == 0 || len(two) == 0 {
		return nil
	}

	if len(two) > len(one) {
		one, two = two, one
	}

	set := make(map[E]struct{})
	for _, value := range one {
		set[value] = struct{}{}
	}

	var result []E
	for _, value := range two {
		if _, exists := set[value]; exists {
			result = append(result, value)
		}
	}
	return result
}
