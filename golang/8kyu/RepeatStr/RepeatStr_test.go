package eightKyu

import "testing"

func TestRepeatStr(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		want := "IIIIII"
		got := RepeatStr(6, "I")
		if want != got {
			t.Errorf("want %q, got %q", want, got)
		}
	})
	t.Run("case 2", func(t *testing.T) {
		want := "HelloHelloHelloHelloHello"
		got := RepeatStr(5, "Hello")
		if want != got {
			t.Errorf("want %q, got %q", want, got)
		}
	})
}
