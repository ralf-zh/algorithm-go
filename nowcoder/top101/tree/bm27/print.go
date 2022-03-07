package bm27

func Print(root *TreeNode) [][]int {
	// write code here
	if root == nil {
		return nil
	}
	var r [][]int

	var q queue
	q.push(root)
	n := 1
	dir := true

	for !q.empty() {
		var n1 int
		var nds []*TreeNode
		for i := 0; i < n; i++ {
			nd := q.pop()
			nds = append(nds, nd)

			if nd.Left != nil {
				q.push(nd.Left)
				n1++
			}
			if nd.Right != nil {
				q.push(nd.Right)
				n1++
			}
		}
		var l []int
		if dir {
			for i := 0; i < len(nds); i++ {
				l = append(l, nds[i].Val)
			}
		} else {
			for i := len(nds) - 1; i >= 0; i-- {
				l = append(l, nds[i].Val)
			}
		}
		dir = !dir
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
