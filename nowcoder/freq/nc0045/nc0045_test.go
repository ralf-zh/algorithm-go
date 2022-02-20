package nc0045

import (
	"reflect"
	"testing"
)

func TestNC0045(t *testing.T) {
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
				Right: &TreeNode{
					Val: 3,
				},
			},
			output: [][]int{{1, 2, 3}, {2, 1, 3}, {2, 3, 1}},
		},
		{
			input:  nil,
			output: [][]int{{}, {}, {}},
		},
	}

	for i, c := range cases {
		expected := threeOrders(c.input)
		if !reflect.DeepEqual(expected, c.output) {
			t.Errorf("[%d] fails, %v != %v", i, expected, c.output)
		}
	}
}
