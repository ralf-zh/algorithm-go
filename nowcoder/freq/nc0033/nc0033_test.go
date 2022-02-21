package nc0033

import (
	"reflect"
	"testing"
)

func TestNC0033(t *testing.T) {
	var cases = []struct {
		input1 *ListNode
		input2 *ListNode
		output []int
	}{
		{
			input1: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
			input2: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 6,
					},
				},
			},
			output: []int{1, 2, 3, 4, 5, 6},
		},
		{
			input1: &ListNode{
				Val: -1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
			input2: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
			output: []int{-1, 1, 2, 3, 4, 4},
		},
	}

	for i, c := range cases {
		chain := Merge(c.input1, c.input2)
		var expected []int
		for chain != nil {
			expected = append(expected, chain.Val)
			chain = chain.Next
		}

		if !reflect.DeepEqual(expected, c.output) {
			t.Errorf("[%d] fails, %v != %v", i, expected, c.output)
		}
	}
}
