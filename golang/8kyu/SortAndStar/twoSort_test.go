package eightKyu

import "testing"

func TestTwoSort(t *testing.T) {
	assertStrings := func(t testing.TB, want, got string) {
		if want != got {
			t.Errorf("want %q, got %q", want, got)
		}
	}
	t.Run("case 1:find lowest", func(t *testing.T) {
		inputs := []string{"Lets", "all", "go", "on", "holiday", "somewhere", "very", "cold"}
		want := "Lets"
		got := FindLowest(inputs)
		assertStrings(t, want, got)
	})
	t.Run("case 2:find lowest", func(t *testing.T) {
		inputs := []string{"turns", "out", "random", "test", "cases", "are", "easier", "than", "writing", "out", "basic", "ones"}
		want := "are"
		got := FindLowest(inputs)
		assertStrings(t, want, got)
	})
	t.Run("case 3:return star", func(t *testing.T) {
		want := "b***i***t***c***o***i***n"
		got := ConvertStarString("bitcoin")
		assertStrings(t, want, got)
	})
	t.Run("case 4: test TwoSort()", func(t *testing.T) {
		want := "a***b***o***u***t"
		got := TwoSort([]string{"lets", "talk", "about", "javascript", "the", "best", "language"})
		assertStrings(t, want, got)
	})
}
