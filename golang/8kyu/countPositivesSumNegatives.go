package eightKyu

func CountPositivesSumNegatives(numbers []int) []int {
	result := make([]int, 0)
	var positive int = 0
	var negative int = 0
	for _, item := range numbers {
		if item > 0 {
			positive += 1
		} else if item < 0 {
			negative += item
		}
	}
	result = append(result, positive)
	result = append(result, negative)
	return result
}
