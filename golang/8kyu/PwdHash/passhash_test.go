package eightKyu

import "testing"

func TestXxx(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		want := "5f4dcc3b5aa765d61d8327deb882cf99"
		got := PassHash("password")
		if want != got {
			t.Errorf("want %q, got %q", want, got)
		}
	})
}
