package nc0050

import (
	"reflect"
	"testing"
)

func TestNC0050(t *testing.T) {
	var cases = []struct {
		input  *ListNode
		k      int
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
			k:      2,
			output: []int{2, 1, 4, 3, 5},
		},
		{
			input:  nil,
			k:      1,
			output: nil,
		},
		{
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			},
			k:      3,
			output: []int{1, 2},
		},
	}

	for i, c := range cases {
		var expected []int
		for r := reverseKGroup(c.input, c.k); r != nil; r = r.Next {
			expected = append(expected, r.Val)
		}
		if !reflect.DeepEqual(expected, c.output) {
			t.Errorf("[%d] fails, %v != %v", i, expected, c.output)
		}
	}
}
