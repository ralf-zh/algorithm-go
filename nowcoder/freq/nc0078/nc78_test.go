package nc0078

import (
	"reflect"
	"testing"
)

/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

func TestNC0078_ReverseList(t *testing.T) {
	var testSuites = []struct {
		input  *ListNode
		output []int
	}{
		{
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
			output: []int{5, 3, 1},
		},
	}

	for i, suite := range testSuites {
		expected := flatList(ReverseList(suite.input))
		if !reflect.DeepEqual(expected, suite.output) {
			t.Errorf("[%d] fails, %v != %v", i, expected, suite.output)
		}
	}
}

func flatList(root *ListNode) []int {
	var r []int
	for root != nil {
		r = append(r, root.Val)
		root = root.Next
	}
	return r
}
