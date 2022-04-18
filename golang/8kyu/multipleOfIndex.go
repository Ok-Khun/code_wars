package eightKyu

func MultipleOfIndex(ints []int) []int {

	result := make([]int, 0)

	for index, item := range ints {
		if index != 0 && item%index == 0 {
			result = append(result, item)
		}
	}

	return result
}
