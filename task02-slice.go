package homework

func reverse(input []int64) (result []int64) {
	result = input
	for i := 0; i < len(input)/2; i++ {
		result[i] = input[len(input)-1-i]
		result[len(input)-1-i] = input[i]
		if len(input)%2 != 0 {
			result[len(input)/2] = input[len(input)/2]
		}

	}
	return
}
