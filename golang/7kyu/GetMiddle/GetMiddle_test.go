package sevenKyu

import (
	"fmt"
	"testing"
)

type testCase struct {
	got  string
	want string
}

func TestGetMiddle(t *testing.T) {

	myTestCases := []testCase{
		{"test", "es"},
		{"testing", "t"},
		{"middle", "dd"},
		{"A", "A"},
	}

	assertEqual := func(t testing.TB, want string, got string) {
		t.Helper()
		if want != got {
			t.Errorf("want %q, got %q", want, got)
		}
	}

	for i, current := range myTestCases {
		currentCase := fmt.Sprintf("case %q", i)
		t.Run(currentCase, func(t *testing.T) {
			want := current.want
			got := GetMiddle(current.got)
			assertEqual(t, want, got)
		})
	}

}
