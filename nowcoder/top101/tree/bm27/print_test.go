package bm27

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
			// {1,2,3,#,#,4,5}
			// 值：[[1],[3,2],[4,5]]
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
			output: [][]int{
				{1},
				{3, 2},
				{4, 5},
			},
		},
		{
			// {8,6,10,5,7,9,11}
			// [[8],[10,6],[5,7,9,11]]
			input: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val: 5,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
				Right: &TreeNode{
					Val: 10,
					Left: &TreeNode{
						Val: 9,
					},
					Right: &TreeNode{
						Val: 11,
					},
				},
			},
			output: [][]int{
				{8},
				{10, 6},
				{5, 7, 9, 11},
			},
		},
		{
			// {1,2,3,4,5}
			// [[1],[3,2],[4,5]]
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			output: [][]int{
				{1},
				{3, 2},
				{4, 5},
			},
		},
	}

	for i, c := range cases {
		expected := Print(c.input)
		if !reflect.DeepEqual(expected, c.output) {
			t.Errorf("[%d] fails, %v != %v", i, expected, c.output)
		}
	}
}
