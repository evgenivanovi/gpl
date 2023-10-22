package slices

func Unique[E comparable](slice []E) bool {

	sliceMap := make(map[E]struct{}, len(slice))

	for _, value := range slice {
		if _, ok := sliceMap[value]; ok {
			return false
		} else {
			sliceMap[value] = struct{}{}
		}
	}

	return true

}
