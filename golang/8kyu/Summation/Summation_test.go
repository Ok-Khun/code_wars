package eightKyu

import "testing"

func TestSummation(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		want := 36
		got := Summation(8)
		if want != got {
			t.Errorf("want %d, got %d", want, got)
		}
	})
}
