package sevenKyu

func GetMiddle(s string) string {
	if len(s) == 0 {
		return ""
	}
	// first check the length of the string
	// if is even or odd
	isEven := len(s)%2 == 0
	mid := len(s) / 2
	if isEven {
		return string(s[mid-1 : mid+1])
	}
	return string(s[mid])
}
