package sevenKyu

import "testing"

func TestHighAndLow(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		want := "5 1"
		got := HighAndLow("1 2 3 4 5")
		if want != got {
			t.Errorf("want %q, got %q", want, got)
		}
	})
	t.Run("case 2", func(t *testing.T) {
		want := "5 -3"
		got := HighAndLow("1 2 -3 4 5")
		if want != got {
			t.Errorf("want %q, got %q", want, got)
		}
	})
}
