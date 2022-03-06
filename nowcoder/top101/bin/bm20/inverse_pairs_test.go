package bm20

import "testing"

func TestInversePairs(t *testing.T) {
	var cases = []struct {
		input  []int
		output int
	}{
		{
			input:  []int{1, 2, 3, 4, 5, 6, 7, 0},
			output: 7,
		},
		{
			input:  []int{1, 2, 3},
			output: 0,
		},
	}

	for i, c := range cases {
		expected := InversePairs(c.input)
		if expected != c.output {
			t.Errorf("[%d] fails,  %v != %v", i, expected, c.output)
		}
	}

}
