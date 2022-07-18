package eightKyu

func SmallestIntegerFinder(arr []int) int {
	current := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < current {
			current = arr[i]
		}
	}
	return current
}
