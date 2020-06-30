package algorithms

import "testing"

func minWindow(s string, t string) string {
	return ""
}

func TestMinWindow(test *testing.T) {
	s := "ADOBECODEBANC"
	t := "ABC"

	expected := "BANC"

	actual := minWindow(s, t)

	if actual != expected {
		test.Errorf("Got %s, want %s", actual, expected)
	}
}
