package main

import (
	eightKyu "codewars/8kyu"
	// sevenKyu "codewars/7kyu"
	"fmt"
)

func main() {
	result := eightKyu.GetGrade(95, 90, 93)
	fmt.Println(result, string(result))
}
