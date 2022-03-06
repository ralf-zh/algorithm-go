package bm02

import (
	"reflect"
	"testing"
)

func TestReverseBetween(t *testing.T) {
	var cases = []struct {
		input  *ListNode
		m      int
		n      int
		output []int
	}{
		{
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
			m:      2,
			n:      4,
			output: []int{1, 4, 3, 2, 5},
		},
		{
			input: &ListNode{
				Val: 5,
			},
			m:      1,
			n:      1,
			output: []int{5},
		},
	}

	for i, c := range cases {
		var expected []int
		for r := reverseBetween(c.input, c.m, c.n); r != nil; r = r.Next {
			expected = append(expected, r.Val)
		}

		if !reflect.DeepEqual(expected, c.output) {
			t.Errorf("[%d] fails, %v != %v", i, expected, c.output)
		}
	}
}
