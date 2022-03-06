package bm04

import (
	"reflect"
	"testing"
)

func TestMergeList(t *testing.T) {
	var cases = []struct {
		input1 *ListNode
		input2 *ListNode
		output []int
	}{
		{
			// {1,3,5},{2,4,6}
			// {1,2,3,4,5,6}
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
			// {},{}
			// {}
			input1: nil,
			input2: nil,
			output: nil,
		},
		{
			// {-1,2,4},{1,3,4}
			// {-1,1,2,3,4,4}
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
		{
			// {1,3,5},{}
			// {1,3,5}
			input1: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
			input2: nil,
			output: []int{1, 3, 5},
		},
	}

	for i, c := range cases {
		var expected []int
		for r := Merge(c.input1, c.input2); r != nil; r = r.Next {
			expected = append(expected, r.Val)
		}
		if !reflect.DeepEqual(expected, c.output) {
			t.Errorf("[%d] fails, %v != %v", i, expected, c.output)
		}
	}
}
