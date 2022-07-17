package eightKyu

import (
	"fmt"
	"testing"
)

type myTestCases struct {
	number int
	result string
}

func TestEvenOrOdd(t *testing.T) {

	testCases := []myTestCases{
		{1, "Odd"},
		{2, "Even"},
		{3, "Odd"},
		{4, "Even"},
		{5, "Odd"},
		{6, "Even"},
		{7, "Odd"},
	}

	for i := 0; i < len(testCases); i++ {
		current := testCases[i]
		myStr := fmt.Sprintf("%d = %q", current.number, current.result)
		t.Run(myStr, func(t *testing.T) {
			want := current.result
			got := EvenOrOdd(current.number)
			if want != got {
				t.Errorf("want %q, got %q", want, got)
			}
		})
	}
}
