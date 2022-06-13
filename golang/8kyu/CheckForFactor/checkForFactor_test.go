package eightKyu

import "testing"

func TestCheckForFactor(t *testing.T) {

	assertBool := func(t testing.TB, got, want bool) {
		t.Helper()
		if want != got {
			t.Errorf("want %t, got %t", want, got)
		}
	}

	t.Run("expects True as a result", func(t *testing.T) {
		want := true
		got := CheckForFactor(10, 2)
		assertBool(t, got, want)
	})

	t.Run("expects False as a result", func(t *testing.T) {
		want := false
		got := CheckForFactor(9, 2)
		assertBool(t, got, want)
	})
}
