package sevenKyu

import (
	// "fmt"
	"math"
)

func Movie(card int, ticket int, perc float64) int {
	count := 1
	systemB := float64(card)
	var systemA int
	for {
		systemA = ticket * count
		systemB += float64(ticket) * nTimes(perc, count)
		if int(math.Ceil(systemB)) < systemA {
			break
		}
		count++
	}
	// fmt.Println(count, systemA, int(math.Ceil(systemB)))
	return count
}

// can implement with math.Powe instead
func nTimes(value float64, count int) float64 {
	result := 1.0

	for i := count; i > 0; i-- {
		result *= value
	}

	return result
}
