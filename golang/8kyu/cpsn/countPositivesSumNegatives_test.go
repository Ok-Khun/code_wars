package eightKyu

import (
	"reflect"
	"testing"
)

func TestCountPositivesSumNegatives(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		arg := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}
		want := []int{10, -65}
		got := CountPositivesSumNegatives(arg)
		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %v, got %v", want, got)
		}
	})
}
