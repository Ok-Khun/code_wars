package eightKyu

func Solution(word string) string {
	// create an empty string
	result := ""
	for i := len(word) - 1; i >= 0; i-- {
		result += string(word[i])
	}
	return result
}
