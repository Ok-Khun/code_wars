package sevenKyu

func isVowel(char rune) bool {
	switch char {
	case 'a':
		return true
	case 'e':
		return true
	case 'i':
		return true
	case 'o':
		return true
	case 'u':
		return true
	}
	return false
}

func GetCount(str string) int {
	result := 0
	for _, char := range str {
		if isVowel(char) {
			result++
		}
	}
	return result
}
