package nc0015

import (
	"reflect"
	"testing"
)

func TestNC0015(t *testing.T) {
	var cases = []struct {
		input  *TreeNode
		output [][]int
	}{
		{
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			output: [][]int{{1}, {2}},
		},
		{
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
			output: [][]int{{1}, {2, 3}, {4, 5}},
		},
	}

	for i, c := range cases {
		expected := levelOrder(c.input)
		if !reflect.DeepEqual(expected, c.output) {
			t.Errorf("[%d] fails, %v != %v", i, expected, c.output)
		}
	}
}
