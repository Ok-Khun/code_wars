package eightKyu

import (
	"reflect"
	"testing"
)

func TestReverseSeq(t *testing.T) {
	t.Run("testing for for ReverseSeq(5)", func(t *testing.T) {
		want := []int{5, 4, 3, 2, 1}
		got := ReverseSeq(5)
		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %v, got %v", want, got)
		}
	})
}
