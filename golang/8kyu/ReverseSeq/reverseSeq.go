package eightKyu

func ReverseSeq(num int) []int {
	// create an empty array
	var result = make([]int, num)
	for i := len(result); i > 0; i-- {
		result[len(result)-i] = i
	}
	return result
}
