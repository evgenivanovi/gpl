package slices

// ContainsAll checks if slice of type E contains all elements of given slice, order independent.
func ContainsAll[E comparable](in []E, need []E) bool {

	inMap := make(map[E]struct{}, len(in))

	for _, key := range in {
		inMap[key] = struct{}{}
	}

	for _, key := range need {
		if _, ok := inMap[key]; !ok {
			return false
		}
	}

	return true

}

// ContainsSame checks if slice of type E contains non-unique elements.
func ContainsSame[E comparable](in []E, need E) bool {
	for key := range in {
		if need == in[key] {
			return false
		}
	}
	return true
}

// ContainsAny checks if slice of type E contains any element from given slice.
func ContainsAny[E comparable](in []E, need []E) bool {
	return len(Intersection(in, need)) > 0
}

// ContainsOne checks if slice of type E contains given element.
func ContainsOne[E comparable](in []E, need E) bool {
	needSlice := make([]E, 1)
	needSlice = append(needSlice, need)
	return ContainsAny(in, needSlice)
}
