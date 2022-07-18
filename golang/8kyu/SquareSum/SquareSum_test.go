package eightKyu

import "testing"

func TestSquareSum(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		want := 9
		got := SquareSum([]int{1, 2, 2})
		if want != got {
			t.Errorf("want %d, got %d", want, got)
		}
	})
}
