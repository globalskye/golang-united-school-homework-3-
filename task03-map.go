package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	keys := make([]int, 0, len(input))
	for v := range input {
		keys = append(keys, v)
	}
	sort.Ints(keys)
	for i, k := range keys {
		result[i] += input[k]
	}

	return
}
