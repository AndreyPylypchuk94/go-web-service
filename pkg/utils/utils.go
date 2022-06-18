package utils

func SliceFilter[T any](source []T, f func(T) bool) []T {
	result := make([]T, 0)

	for _, el := range source {
		if f(el) {
			result = append(result, el)
		}
	}

	return result
}
