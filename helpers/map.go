package helpers

func FlatToMap[T comparable](flat []T) map[T]T {
	var null T
	result := make(map[T]T)

	n := len(flat)
	k := n / 2

	for i := 0; i < k; i++ {
		result[flat[2*i]] = flat[2*i+1]
	}

	if n%2 != 0 {
		result[flat[2*k]] = null
	}

	return result
}

func MapToFlat[T comparable](input map[T]T, merge map[T]T) []T {
	flat := []T{}

	for key, value := range input {
		output, ok := merge[key]

		flat = append(flat, key)

		if !ok {
			flat = append(flat, value)
		} else {
			flat = append(flat, output)
		}
	}

	return flat
}
