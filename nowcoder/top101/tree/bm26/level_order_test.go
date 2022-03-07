package bm26

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	var cases = []struct {
		input  *TreeNode
		output [][]int
	}{
		{
			// 1,2
			// [[1], [2]]
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			output: [][]int{
				{1},
				{2},
			},
		},
		{
			// {1,2,3,4,#,#,5}
			// [[1],[2,3],[4,5]]
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
			output: [][]int{
				{1},
				{2, 3},
				{4, 5},
			},
		},
	}

	for i, c := range cases {
		expected := levelOrder(c.input)
		if !reflect.DeepEqual(expected, c.output) {
			t.Errorf("[%d] fails, %v != %v", i, expected, c.output)
		}
	}
}
