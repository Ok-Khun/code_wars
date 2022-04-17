package eightKyu

import (
	"strings"
)

// you can easily solve this with sort package!

func TwoSort(arr []string) string {

	var smallestIndex int = 0

	for i := 1; i < len(arr); i++ {
		item := arr[i]
		curr := arr[smallestIndex]
		if item[0] < curr[0] {
			smallestIndex = i
		} else if item[0] == curr[0] {
			if item[1] < curr[1] {
				smallestIndex = i
			}
		}
	}

	splitStr := strings.Split(arr[smallestIndex], "")

	var result string

	for i := 0; i < len(splitStr)-1; i++ {
		result += splitStr[i] + "***"
	}

	result += splitStr[len(splitStr)-1]

	return result
}
