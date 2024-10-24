package slice

import "reflect"

func AllSameType(slice []any) (bool, reflect.Type) {
	if len(slice) <= 1 {
		return true, nil
	}

	TYPE := reflect.TypeOf(slice[0])

	for _, value := range slice[1:] {
		if reflect.TypeOf(value) != TYPE {
			return false, nil
		}
	}

	return true, TYPE
}
