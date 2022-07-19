package eightKyu

import "testing"

func TestStringToNumber(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		want := 1234
		got := StringToNumber("1234")
		if want != got {
			// throw error
			t.Errorf("want %T, got %T", want, got)
		}
	})
}
