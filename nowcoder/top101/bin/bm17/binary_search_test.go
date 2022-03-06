package bm17

import "testing"

func TestBinarySearch(t *testing.T) {
	var cases = []struct {
		input  []int
		target int
		ouput  int
	}{
		{
			input:  []int{-1, 0, 3, 4, 6, 10, 13, 14},
			target: 13,
			ouput:  6,
		},
		{
			input:  nil,
			target: 3,
			ouput:  -1,
		},
		{
			input:  []int{-1, 0, 3, 4, 6, 10, 13, 14},
			target: 2,
			ouput:  -1,
		},
		{
			input:  []int{9999},
			target: 9999,
			ouput:  0,
		},
	}

	for i, c := range cases {
		expected := search(c.input, c.target)
		if expected != c.ouput {
			t.Errorf("[%d] fails, %v != %v", i, expected, c.ouput)
		}
	}
}
