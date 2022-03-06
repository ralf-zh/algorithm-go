package bm19

import "testing"

func TestFindPeakElement(t *testing.T) {
	var cases = []struct {
		input  []int
		output int
	}{
		{
			input:  []int{2, 4, 1, 2, 7, 8, 4},
			output: 5,
		},
		{
			input:  []int{1, 2, 3, 1},
			output: 2,
		},
	}

	for i, c := range cases {
		expected := findPeakElement(c.input)
		if expected != c.output {
			t.Errorf("[%d] fails, %v != %v", i, expected, c.output)
		}
	}
}
