package eightKyu

func EvenOrOdd(number int) string {
	result := number % 2
	if result == 0 {
		return "Even"
	}
	return "Odd"
}
