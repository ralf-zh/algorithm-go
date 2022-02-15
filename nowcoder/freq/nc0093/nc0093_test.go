package nc0093

import (
	"reflect"
	"testing"
)

func TestLRU(t *testing.T) {
	var testSuites = []struct {
		operators [][]int
		k         int
		expect    []int
	}{
		{
			operators: [][]int{
				{1, 1, 1}, {1, 2, 2}, {1, 3, 2}, {2, 1}, {1, 4, 4}, {2, 2},
			},
			k:      3,
			expect: []int{1, -1},
		},
	}

	for i, suite := range testSuites {
		expected := LRU(suite.operators, suite.k)
		if !reflect.DeepEqual(expected, suite.expect) {
			t.Errorf("[%d] fails, %v != %v", i, expected, suite.expect)
		}
	}
}
