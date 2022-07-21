package sevenKyu

import (
	"fmt"
	"strconv"
	"strings"
)

func convertStringToInt(arg string) int {
	result, _ := strconv.Atoi(arg)
	return result
}

func HighAndLow(str string) string {
	// split the string by space
	strArr := strings.Split(str, " ") // this is an array of string
	first := strArr[0]
	low := convertStringToInt(first)
	high := convertStringToInt(first)

	for _, current := range strArr {
		if low < convertStringToInt(current) {
			low = convertStringToInt(current)
		}
		if high > convertStringToInt(current) {
			high = convertStringToInt(current)
		}
	}

	result := fmt.Sprintf("%d %d", low, high)

	return result
}
