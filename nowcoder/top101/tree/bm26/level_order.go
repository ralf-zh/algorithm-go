package bm26

func levelOrder(root *TreeNode) [][]int {
	// write code here
	if root == nil {
		return nil
	}
	var r [][]int

	var q queue
	q.push(root)
	n := 1

	for !q.empty() {
		var l []int
		var n1 int
		for i := 0; i < n; i++ {
			nd := q.pop()
			l = append(l, nd.Val)

			if nd.Left != nil {
				q.push(nd.Left)
				n1++
			}
			if nd.Right != nil {
				q.push(nd.Right)
				n1++
			}
		}
		r = append(r, l)
		n = n1
	}
	return r
}

type queue []*TreeNode

func (q *queue) empty() bool {
	return q == nil || len(*q) == 0
}

func (q *queue) push(nd *TreeNode) {
	*q = append(*q, nd)
}

func (q *queue) pop() *TreeNode {
	nd := (*q)[0]
	*q = (*q)[1:]
	return nd
}
