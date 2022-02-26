package nc0022

import (
	"reflect"
	"testing"
)

func TestNC0022(t *testing.T) {
	var cases = []struct {
		A      []int
		m      int
		B      []int
		n      int
		output []int
	}{
		{
			A:      []int{4, 5, 6, 0, 0, 0},
			m:      3,
			B:      []int{1, 2, 3},
			n:      3,
			output: []int{1, 2, 3, 4, 5, 6},
		},
		{
			A:      []int{1, 2, 3, 0, 0, 0},
			m:      3,
			B:      []int{2, 5, 6},
			n:      3,
			output: []int{1, 2, 2, 3, 5, 6},
		},
	}

	for i, c := range cases {
		merge(c.A, c.m, c.B, c.n)
		if !reflect.DeepEqual(c.A, c.output) {
			t.Errorf("[%d] fails, %v != %v", i, c.A, c.output)
		}
	}
}
