package lc0001

import (
	"reflect"
	"testing"
)

func Test_LC0001(t *testing.T) {
	var testSuite = []struct {
		nums   []int
		target int
		expect []int
	}{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			expect: []int{0, 1},
		},
		{
			nums:   []int{3, 2, 4},
			target: 6,
			expect: []int{1, 2},
		},
		{
			nums:   []int{3, 3},
			target: 6,
			expect: []int{0, 1},
		},
	}

	for i, suite := range testSuite {
		expected := twoSum(suite.nums, suite.target)

		if !reflect.DeepEqual(expected, suite.expect) {
			t.Errorf("[%d] failed", i)
		}
	}
}
