package eightKyu

func CheckForFactor(base int, factor int) bool {
	remainder := base % factor
	return remainder == 0
}
