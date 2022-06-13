package eightKyu

import "testing"

func TestCountSheep(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		want := "1 sheep...2 sheep...3 sheep..."
		got := CountSheep(3)
		if want != got {
			t.Errorf("want %q, got %q", want, got)
		}
	})
}
