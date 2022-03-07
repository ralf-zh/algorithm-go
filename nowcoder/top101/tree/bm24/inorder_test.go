package bm24

import (
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	var cases = []struct {
		input  *TreeNode
		output []int
	}{
		{
			// {1,2,#,#,3}
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			output: []int{2, 3, 1},
		},
		{
			// {}
			input:  nil,
			output: nil,
		},
		{
			// {1, 2}
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			output: []int{2, 1},
		},
		{
			// {1,#,2}
			input: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
			output: []int{1, 2},
		},
	}

	for i, c := range cases {
		expected := inorderTraversal(c.input)
		if !reflect.DeepEqual(expected, c.output) {
			t.Errorf("[%d] fails, %v != %v", i, expected, c.output)
		}
	}
}
