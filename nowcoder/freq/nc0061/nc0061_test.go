package nc0061

import (
	"reflect"
	"testing"
)

func TestNC0061(t *testing.T) {
	var testSuite = []struct {
		nums   []int
		target int
		expect []int
	}{
		{
			nums:   []int{3, 2, 4},
			target: 6,
			expect: []int{2, 3},
		},
		{
			nums:   []int{20, 70, 110, 150},
			target: 90,
			expect: []int{1, 2},
		},
	}

	for i, suite := range testSuite {
		expected := twoSum(suite.nums, suite.target)

		if !reflect.DeepEqual(expected, suite.expect) {
			t.Errorf("[%d] failed, %v != %v", i, expected, suite.expect)
		}
	}
}
