package eightKyu

func ReverseSeq(n int) []int {
	arr := make([]int, 0)
	for i := n; i > 0; i-- {
		arr = append(arr, i)
	}
	return arr
}
