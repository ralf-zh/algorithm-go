package bm28

func maxDepth(root *TreeNode) int {
	// write code here
	var depth func(root *TreeNode) int
	depth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := depth(root.Left)
		r := depth(root.Right)

		if r > l {
			l = r
		}
		return l + 1
	}

	return depth(root)
}
