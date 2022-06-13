package eightKyu

import "testing"

func TestGetGrade(t *testing.T) {

	assertRunes := func(t testing.TB, want, got rune) {
		t.Helper()
		if want != got {
			t.Errorf("want %q, got %q", want, got)
		}
	}
	t.Run("expects 'A' as a result", func(t *testing.T) {
		want := 'A'
		got := GetGrade(95, 90, 93)
		assertRunes(t, want, got)
	})
	t.Run("expects 'B' as a result", func(t *testing.T) {
		want := 'B'
		got := GetGrade(70, 70, 100)
		assertRunes(t, want, got)
	})
	t.Run("expects 'C' as a result", func(t *testing.T) {
		want := 'C'
		got := GetGrade(70, 70, 70)
		assertRunes(t, want, got)
	})
	t.Run("expects 'D' as a result", func(t *testing.T) {
		want := 'D'
		got := GetGrade(65, 70, 59)
		assertRunes(t, want, got)
	})
	t.Run("expects 'F' as a result", func(t *testing.T) {
		want := 'F'
		got := GetGrade(44, 55, 52)
		assertRunes(t, want, got)
	})
}
