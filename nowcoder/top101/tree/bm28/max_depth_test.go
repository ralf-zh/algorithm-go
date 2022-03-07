package bm28

import "testing"

func TestMaxDepth(t *testing.T) {
	var cases = []struct {
		input  *TreeNode
		output int
	}{
		{
			// {1, 2}
			// 2
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			output: 2,
		},
		{
			// {1,2,3,4,#,#,5}
			// 3
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
			output: 3,
		},
	}

	for i, c := range cases {
		expected := maxDepth(c.input)
		if expected != c.output {
			t.Errorf("[%d] fails, %v != %v", i, expected, c.output)
		}
	}
}
