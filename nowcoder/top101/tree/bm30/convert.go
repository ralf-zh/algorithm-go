package bm30

func Convert(pRootOfTree *TreeNode) *TreeNode {
	// write code here
	var prev, head *TreeNode

	var stk = make(stack, 0, 100000)
	var nd = pRootOfTree

	for !stk.empty() || nd != nil { // TODO
		if nd != nil {
			stk.push(nd)
			nd = nd.Left
		} else {
			nd = stk.pop()

			// visitor begin
			nd.Left = prev
			if prev == nil {
				head = nd
			} else {
				prev.Right = nd
			}
			prev = nd
			// visitor end

			nd = nd.Right
		}
	}
	prev.Right = head

	return head
}

type stack []*TreeNode

func (s *stack) empty() bool {
	return s == nil || len(*s) == 0
}

func (s *stack) push(nd *TreeNode) {
	*s = append(*s, nd)
}

func (s *stack) pop() *TreeNode {
	nd := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return nd
}

func Convert2(pRootOfTree *TreeNode) *TreeNode {
	// write code here
	var prev, head *TreeNode
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)

		root.Left = prev
		if prev != nil {
			prev.Right = root
		} else {
			head = root
		}
		prev = root

		traverse(root.Right)
	}
	traverse(pRootOfTree)
	prev.Right = head

	return head
}

func Convert1(pRootOfTree *TreeNode) *TreeNode {
	// write code here
	var l []*TreeNode
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)
		l = append(l, root)
		traverse(root.Right)
	}
	traverse(pRootOfTree)
	if len(l) == 0 {
		return nil
	}

	for i := 0; i < len(l); i++ {
		cur := l[i]
		prev := l[(i+len(l)-1)%len(l)]
		next := l[(i+len(l)+1)%len(l)]

		cur.Left = prev
		cur.Right = next
	}
	return l[0]
}
