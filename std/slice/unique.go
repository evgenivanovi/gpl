package slice

func Unique[E comparable](slice []E) bool {
	set := make(map[E]struct{}, len(slice))

	for _, elem := range slice {
		if _, ok := set[elem]; ok {
			return false
		} else {
			set[elem] = struct{}{}
		}
	}

	return true
}
