package bm18

import "testing"

func TestFind(t *testing.T) {
	var cases = []struct {
		input  [][]int
		target int
		output bool
	}{
		{
			input: [][]int{
				{1, 2, 8, 9},
				{2, 4, 9, 12},
				{4, 7, 10, 13},
				{6, 8, 11, 15},
			},
			target: 7,
			output: true,
		},
		{
			input: [][]int{
				{2},
			},
			target: 1,
			output: false,
		},
		{
			input: [][]int{
				{1, 2, 8, 9},
				{2, 4, 9, 12},
				{4, 7, 10, 13},
				{6, 8, 11, 15},
			},
			target: 3,
			output: false,
		},
	}

	for i, c := range cases {
		expected := Find(c.target, c.input)
		if expected != c.output {
			t.Errorf("[%d] fails, %v != %v", i, expected, c.output)
		}
	}
}
