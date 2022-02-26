package nc0041

import "testing"

func TestNC0041(t *testing.T) {
	var cases = []struct {
		input  []int
		output int
	}{
		{
			input:  []int{2, 3, 4, 5},
			output: 4,
		},
		{
			input:  []int{2, 2, 3, 4, 3},
			output: 3,
		},
		{
			input:  []int{9},
			output: 1,
		},
		{
			input:  []int{1, 2, 3, 1, 2, 3, 2, 2},
			output: 3,
		},
		{
			input:  []int{2, 2, 3, 4, 8, 99, 3},
			output: 5,
		},
		{
			input:  []int{3, 3, 2, 1, 3, 3, 3, 1},
			output: 3,
		},
	}

	for i, c := range cases {
		expected := maxLength(c.input)
		if expected != c.output {
			t.Errorf("[%d] fails, %v != %v", i, expected, c.output)
		}
	}
}
