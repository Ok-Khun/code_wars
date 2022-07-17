package eightKyu

func IsPositive(number int) bool {
	return number >= 0
}

func PositiveSum(numbers []int) int {
	sum := 0

	if len(numbers) == 0 {
		return sum
	}

	for i := 0; i < len(numbers); i++ {
		current := numbers[i]
		if IsPositive(current) {
			sum += current
		}
	}

	return sum
}
