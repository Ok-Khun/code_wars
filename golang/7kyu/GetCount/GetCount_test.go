package sevenKyu

import "testing"

func TestGetCount(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		want := 5
		got := GetCount("abracadabra")
		if want != got {
			t.Errorf("want %d, got %d", want, got)
		}
	})
}
