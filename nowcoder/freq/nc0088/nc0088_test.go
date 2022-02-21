package nc0088

import (
	"testing"
)

func TestNC0088(t *testing.T) {
	var cases = []struct {
		input  []int
		n      int
		K      int
		output int
	}{
		{
			input:  []int{1, 3, 5, 2, 2},
			n:      5,
			K:      3,
			output: 2,
		},
		{
			input:  []int{10, 10, 9, 9, 8, 7, 5, 6, 4, 3, 4, 2},
			n:      12,
			K:      3,
			output: 9,
		},
	}

	for i, c := range cases {
		expected := findKth(c.input, c.n, c.K)
		if expected != c.output {
			t.Errorf("[%d] fails, %v != %v", i, expected, c.output)
		}
	}
}
