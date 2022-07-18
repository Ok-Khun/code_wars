package eightKyu

func SquareSum(arr []int) int {
	result := 0
	for _, num := range arr {
		sq := num * num
		result += sq
	}
	return result
}
