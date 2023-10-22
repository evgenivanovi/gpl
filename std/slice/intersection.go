package slices

// Intersection returns intersection for slices of various built-in types.
func Intersection[E comparable](one []E, two []E) []E {

	if len(one) == 0 || len(two) == 0 {
		return nil
	}

	first, second := one, two
	if len(two) > len(one) {
		first, second = two, one
	}

	firstMap := make(map[E]struct{})
	for _, i := range first {
		firstMap[i] = struct{}{}
	}

	var result []E
	for _, value := range second {
		if _, exists := firstMap[value]; exists {
			result = append(result, value)
		}
	}
	return result

}
