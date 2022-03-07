package bm23

import (
	"reflect"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	var cases = []struct {
		input  *TreeNode
		output []int
	}{
		{
			// {1,#,2,3}
			input: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			output: []int{1, 2, 3},
		},
	}

	for i, c := range cases {
		expected := preorderTraversal(c.input)
		if !reflect.DeepEqual(expected, c.output) {
			t.Errorf("[%d] fails, %v != %v", i, expected, c.output)
		}
	}
}
