package lc0026

import "testing"

func Test_LC0026(t *testing.T) {
	var testSuites = []struct {
		nums   []int
		expect int
	}{
		{
			nums:   []int{1, 1, 2},
			expect: 2,
		},
		{
			nums:   []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expect: 5,
		},
	}
	for i, suite := range testSuites {
		expected := removeDuplicates(suite.nums)
		if expected != suite.expect {
			t.Errorf("[%d] fails", i)
		}
	}
}
