package nc0053

import (
	"reflect"
	"testing"
)

func TestNC0053(t *testing.T) {
	var cases = []struct {
		input  *ListNode
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
			n:      2,
			output: []int{1, 2, 3, 5},
		},
		{
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			},
			n:      2,
			output: []int{2},
		},
	}

	for i, c := range cases {
		var expected []int
		for r := removeNthFromEnd(c.input, c.n); r != nil; r = r.Next {
			expected = append(expected, r.Val)
		}

		if !reflect.DeepEqual(expected, c.output) {
			t.Errorf("[%d] fails, %v != %v", i, expected, c.output)
		}

	}
}
