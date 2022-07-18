package eightKyu

import "testing"

func TestSmallestIntegerFinder(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {
		arr := []int{34, 15, 88, 2}
		want := 2
		got := SmallestIntegerFinder(arr)
		if want != got {
			t.Errorf("want %d, got %d", want, got)
		}
	})

}
