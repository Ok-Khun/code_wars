package eightKyu

import "testing"

func TestPositiveSum(t *testing.T) {
	assertEqual := func(want int, got int, t testing.TB) {
		t.Helper()
		if want != got {
			t.Errorf("want %d, got %d", want, got)
		}
	}
	t.Run("case 1", func(t *testing.T) {
		want := 15
		got := PositiveSum([]int{1, 2, 3, 4, 5})
		assertEqual(want, got, t)
	})

	t.Run("case 2", func(t *testing.T) {
		want := 13
		got := PositiveSum([]int{1, -2, 3, 4, 5})
		assertEqual(want, got, t)
	})

}
