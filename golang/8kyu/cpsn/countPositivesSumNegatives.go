package eightKyu

func CountPositivesSumNegatives(numbers []int) []int {
	var result = make([]int, 2)
	var countPos int = 0
	var sumOfNeg int = 0
	for _, i := range numbers {
		if i > 0 {
			countPos++
		} else if i < 0 {
			sumOfNeg += i
		}
	}
	result[0] = countPos
	result[1] = sumOfNeg
	return result
}
