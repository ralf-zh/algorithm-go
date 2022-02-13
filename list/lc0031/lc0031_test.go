package lc0031

import (
	"reflect"
	"testing"
)

func Test_LC0031(t *testing.T) {
	var testSuites = []struct {
		nums   []int
		expect []int
	}{
		{
			nums:   []int{1, 2, 3},
			expect: []int{1, 3, 2},
		},
		{
			nums:   []int{3, 2, 1},
			expect: []int{1, 2, 3},
		},
		{
			nums:   []int{1, 1, 5},
			expect: []int{1, 5, 1},
		},
		{
			nums:   []int{1, 3, 2},
			expect: []int{2, 1, 3},
		},
		{
			nums:   []int{4, 2, 0, 2, 3, 2, 0},
			expect: []int{4, 2, 0, 3, 0, 2, 2},
		},
	}

	for i, suite := range testSuites {
		nextPermutation(suite.nums)
		if !reflect.DeepEqual(suite.nums, suite.expect) {
			t.Errorf("[%d] fails, %v != %v", i, suite.nums, suite.expect)
		}
	}
}
