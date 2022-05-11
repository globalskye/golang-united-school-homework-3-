package homework

func reverse(input []int64) (result []int64) {
	result = make([]int64, len(input))
	for i := 0; i < len(input)/2; i++ {
		result[i] = input[len(input)-i-1]
		result[len(input)-i-1] = input[i]
		if len(input)%2 != 0 {
			result[len(input)/2] = input[len(input)/2]
		}

	}
	return
}
