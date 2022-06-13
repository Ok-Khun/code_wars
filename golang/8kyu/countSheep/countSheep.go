package eightKyu

import (
	"strconv"
)

func CountSheep(n int) string {
	sheep := " sheep..."
	result := ""
	for i := 1; i <= n; i++ {
		str := strconv.Itoa(i)
		result += str + sheep
	}
	return result
}
