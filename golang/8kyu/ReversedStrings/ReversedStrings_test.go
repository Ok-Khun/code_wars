package eightKyu

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {

	cases := [][2]string{
		{"world", "dlrow"},
		{"goLang", "gnaLog"},
	}

	for i := 0; i < len(cases); i++ {
		current := cases[i]
		myStr := fmt.Sprintf("%q, %q", current[0], current[1])
		t.Run(myStr, func(t *testing.T) {
			want := current[1]
			got := Solution(current[0])
			if want != got {
				t.Errorf("want %q, got %q", want, got)
			}
		})

	}

}
