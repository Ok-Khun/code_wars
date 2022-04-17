package eightKyu

import "strconv"

func CountSheep(num int) string {
	var curr int = 1
	var result string
	for curr <= num {
		result += strconv.Itoa(curr) + " sheep..."
		curr += 1
	}
	return result
}
