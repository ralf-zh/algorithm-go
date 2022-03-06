package bm01

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	var cases = []struct {
		input  *ListNode
		output []int
	}{
		{
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
			output: []int{3, 2, 1},
		},
		{
			input:  nil,
			output: nil,
		},
	}

	for i, c := range cases {
		var expected []int
		for r := ReverseList(c.input); r != nil; r = r.Next {
			expected = append(expected, r.Val)
		}
		if !reflect.DeepEqual(expected, c.output) {
			t.Errorf("[%d] fails, %v != %v", i, expected, c.output)
		}
	}
}
