package lc0027

import "testing"

func Test_LC0027(t *testing.T) {
	var testSuites = []struct {
		nums   []int
		val    int
		expect int
	}{
		{
			nums:   []int{3, 2, 2, 3},
			val:    3,
			expect: 2,
		},
		{
			nums:   []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:    2,
			expect: 5,
		},
	}
	for i, suite := range testSuites {
		expected := removeElement(suite.nums, suite.val)
		if expected != suite.expect {
			t.Errorf("[%d] fails, %d != %d", i, expected, suite.expect)
		}
	}
}
