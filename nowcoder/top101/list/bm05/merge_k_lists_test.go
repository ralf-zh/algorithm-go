package bm05

import (
	"reflect"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	var cases = []struct {
		input  [][]int
		output []int
	}{
		{
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6, 7},
			},
			output: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			input: [][]int{
				{1, 2},
				{1, 4, 5},
				{6},
			},
			output: []int{1, 1, 2, 4, 5, 6},
		},
	}

	for i, c := range cases {
		var expected []int
		l := makeKLists(c.input)

		for r := mergeKLists(l); r != nil; r = r.Next {
			expected = append(expected, r.Val)
		}
		if !reflect.DeepEqual(expected, c.output) {
			t.Errorf("[%d] fails, %v != %v", i, expected, c.output)
		}
	}
}

func makeKLists(input [][]int) []*ListNode {
	var l []*ListNode
	for _, nums := range input {
		l = append(l, makeList(nums))
	}
	return l
}

func makeList(input []int) *ListNode {
	var l, t *ListNode
	for _, n := range input {
		if l == nil {
			l = &ListNode{Val: n}
			t = l
		} else {
			t.Next = &ListNode{Val: n}
			t = t.Next
		}
	}
	return l
}
