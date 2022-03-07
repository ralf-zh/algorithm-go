package bm29

import "testing"

func TestHasPathSum(t *testing.T) {
	var cases = []struct {
		input  *TreeNode
		sum    int
		output bool
	}{
		{
			// {5,4,8,1,11,#,9,#,#,2,7},22
			// true
			input: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 11,
						Left: &TreeNode{
							Val: 2,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Right: &TreeNode{
						Val: 9,
					},
				},
			},
			sum:    22,
			output: true,
		},
		{
			// {1,2},0
			// false
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			sum:    0,
			output: false,
		},
		{
			// {1,2},3
			// true
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			sum:    3,
			output: true,
		},
		{
			// {},0
			// false
			input:  nil,
			sum:    0,
			output: false,
		},
	}

	for i, c := range cases {
		expected := hasPathSum(c.input, c.sum)
		if expected != c.output {
			t.Errorf("[%d] fails, %v != %v", i, expected, c.output)
		}
	}
}
