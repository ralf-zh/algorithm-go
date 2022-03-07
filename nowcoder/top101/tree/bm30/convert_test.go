package bm30

import (
	"reflect"
	"testing"
)

func TestConvert(t *testing.T) {
	var cases = []struct {
		input  *TreeNode
		output []int
	}{
		{
			// {10,6,14,4,8,12,16}
			// [4, 6, 8, 10, 12, 14, 16]
			input: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 8,
					},
				},
				Right: &TreeNode{
					Val: 14,
					Left: &TreeNode{
						Val: 12,
					},
					Right: &TreeNode{
						Val: 16,
					},
				},
			},
			output: []int{4, 6, 8, 10, 12, 14, 16},
		},
		{
			// {5,4,#,3,#,2,#,1}
			// [1,2,3,4,5]
			input: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 2,
							Left: &TreeNode{
								Val: 1,
							},
						},
					},
				},
			},
			output: []int{1, 2, 3, 4, 5},
		},
	}

	for i, c := range cases {
		var expeced []int
		r := Convert(c.input)
		for nd, meet := r, false; nd != r || !meet; nd, meet = nd.Right, true {
			expeced = append(expeced, nd.Val)
		}

		if !reflect.DeepEqual(expeced, c.output) {
			t.Errorf("[%d] fails, %v != %v", i, expeced, c.output)
		}
	}
}
