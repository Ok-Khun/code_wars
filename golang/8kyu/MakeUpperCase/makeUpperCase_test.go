package eightKyu

import "testing"

func TestMakeUpperCase(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		want := "HELLO"
		got := MakeUpperCase("hello")
		if want != got {
			t.Errorf("want %q, got %q", want, got)
		}
	})
}
