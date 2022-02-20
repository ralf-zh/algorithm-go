package nc0015

/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 *
 * @param root TreeNode类
 * @return int整型二维数组
 */
func levelOrder(root *TreeNode) [][]int {
	// write code here

	var r [][]int
	var queue TreeNodeQueue

	queue.Push(root)
	l := 1
	for !queue.Empty() {
		var (
			width = 0
			vals  []int
		)

		for i := 0; i < l; i++ {
			n := queue.Pop()
			if n != nil {
				vals = append(vals, n.Val)

				if n.Left != nil {
					queue.Push(n.Left)
					width++
				}
				if n.Right != nil {
					queue.Push(n.Right)
					width++
				}
			}
		}
		if len(vals) == 0 {
			break
		}

		l = width
		r = append(r, vals)
	}

	return r
}

type TreeNodeQueue []*TreeNode

func (q *TreeNodeQueue) Push(n *TreeNode) {
	(*q) = append((*q), n)
}

func (q *TreeNodeQueue) Pop() *TreeNode {
	n := (*q)[0]
	(*q) = (*q)[1:]

	return n
}

func (q *TreeNodeQueue) Empty() bool {
	return q == nil || len(*q) == 0
}
