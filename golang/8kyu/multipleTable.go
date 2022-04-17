package eightKyu

import "strconv"

func MultipleTable(number int) string {
	var result string
	for i := 1; i <= 10; i++ {
		p := i * number
		result += strconv.Itoa(i) + " * " + strconv.Itoa(number) + " = " + strconv.Itoa(p)
		if i != 10 {
			result += "\n"
		}
	}
	return result
}
