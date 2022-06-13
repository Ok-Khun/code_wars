package eightKyu

import (
	"sort"
)

func FindLowest(arr []string) string {
	sort.Strings(arr)
	return arr[0]
}

func ConvertStarString(str string) string {
	result := ""
	stars := "***"
	for i := 0; i < len(str)-1; i++ {
		result += string(str[i]) + stars
	}
	result += string(str[len(str)-1])
	return result
}

func TwoSort(arr []string) string {
	arrCopied := arr
	return ConvertStarString(FindLowest(arrCopied))
}
