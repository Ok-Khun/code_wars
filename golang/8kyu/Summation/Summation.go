package eightKyu

func Summation(num int) int {
	if num == 0 {
		return 0
	}
	result := 0
	current := num
	for current > 0 {
		result += current
		current--
	}
	return result
}
