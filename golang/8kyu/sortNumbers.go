package eightKyu

// selection sort
// note: test in codewar not working.
func SortNumbers(numbers []int) []int {
	numArr := numbers
	resultArr := make([]int, 0)
	for len(numArr) > 0 {
		i, n := findSmallest(numArr)
		resultArr = append(resultArr, n)
		numArr = append(numArr[:i], numArr[i+1:]...)
	}
	return resultArr
}

func findSmallest(numbers []int) (int, int) {
	result := numbers[0]
	resultIndex := 0
	for index, item := range numbers {
		if item < result {
			result = item
			resultIndex = index
		}
	}
	return resultIndex, result
}
