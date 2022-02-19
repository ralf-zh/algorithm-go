package lc0035

import (
	"reflect"
	"testing"
)

func Test_LC0035(t *testing.T) {
	var testSuites = []struct {
		nums   []int
		target int
		expect int
	}{
		{
			nums:   []int{1, 3, 5, 6},
			target: 5,
			expect: 2,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 2,
			expect: 1,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 7,
			expect: 4,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 4,
			expect: 2,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 0,
			expect: 0,
		},
	}

	for i, suite := range testSuites {
		expected := searchInsert(suite.nums, suite.target)
		if !reflect.DeepEqual(expected, suite.expect) {
			t.Errorf("[%d] fails, %v != %v", i, expected, suite.expect)
		}
	}
}
