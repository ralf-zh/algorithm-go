package bm21

import "testing"

func TestMinNumberInRotateArrray(t *testing.T) {
	var cases = []struct {
		input  []int
		output int
	}{
		{
			input:  []int{3, 4, 5, 1, 2},
			output: 1,
		},
		{
			input:  []int{3, 100, 200, 3},
			output: 3,
		},
		{
			input:  []int{1, 0, 1, 1, 1},
			output: 0,
		},
		{
			input:  []int{2, 2, 2, 1, 2},
			output: 1,
		},
	}

	for i, c := range cases {
		expeced := minNumberInRotateArray(c.input)
		if expeced != c.output {
			t.Errorf("[%d] fails, %v != %v", i, expeced, c.output)
		}
	}
}
