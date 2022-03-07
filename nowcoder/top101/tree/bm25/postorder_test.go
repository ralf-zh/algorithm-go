package bm25

import (
	"reflect"
	"testing"
)

func TestPostorderTraversal(t *testing.T) {
	var cases = []struct {
		input  *TreeNode
		output []int
	}{
		{
			// {1,#,2, 3}
			input: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			output: []int{3, 2, 1},
		},
		{
			// {1}
			input: &TreeNode{
				Val: 1,
			},
			output: []int{1},
		},
	}

	for i, c := range cases {
		expected := postorderTraversal(c.input)
		if !reflect.DeepEqual(expected, c.output) {
			t.Errorf("[%d] fails, %v != %v", i, expected, c.output)
		}
	}
}
