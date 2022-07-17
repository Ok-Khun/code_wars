package eightKyu

import (
	"fmt"
	"testing"
)

// simple test case
type testCase struct {
	numbers [2]int
	result  int
}

func TestMultiply(t *testing.T) {

	myCases := []testCase{
		{[2]int{2, 2}, 4},
		{[2]int{5, 2}, 10},
		{[2]int{3, 8}, 24},
	}

	for i := 0; i < len(myCases); i++ {
		current := myCases[i]
		myStr := fmt.Sprintf("%d x %d = %d", current.result, current.numbers[0], current.numbers[1])
		t.Run(myStr, func(t *testing.T) {
			want := current.result
			got := Multiply(current.numbers[0], current.numbers[1])
			if got != want {
				t.Errorf("want %d, got %d", want, got)
			}
		})
	}
}
