package nc0019

import "testing"

func TestNC0019(t *testing.T) {
	var cases = []struct {
		input  []int
		output int
	}{
		{
			input:  []int{1, -2, 3, 10, -4, 7, 2, -5},
			output: 18,
		},
		{
			input:  []int{2},
			output: 2,
		},
		{
			input:  []int{-10},
			output: -10,
		},
	}

	for i, c := range cases {
		expected := FindGreatestSumOfSubArray(c.input)
		if expected != c.output {
			t.Errorf("[%d] %v != %v", i, expected, c.output)
		}
	}
}
