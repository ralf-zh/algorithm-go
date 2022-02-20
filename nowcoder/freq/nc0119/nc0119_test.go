package nc0119

import (
	"reflect"
	"testing"
)

func TestNC0119(t *testing.T) {
	var cases = []struct {
		input  []int
		k      int
		output []int
	}{
		{
			input:  []int{4, 5, 1, 6, 2, 7, 3, 8},
			k:      4,
			output: []int{1, 2, 3, 4},
		},
		{
			input:  []int{1},
			k:      0,
			output: []int{},
		},
		{
			input:  []int{0, 1, 2, 1, 2},
			k:      3,
			output: []int{0, 1, 1},
		},
	}

	for i, c := range cases {
		expected := GetLeastNumbers_Solution(c.input, c.k)
		if !reflect.DeepEqual(expected, c.output) {
			t.Errorf("[%d] fails, %v != %v", i, expected, c.output)
		}
	}
}
