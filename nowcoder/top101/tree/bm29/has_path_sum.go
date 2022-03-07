package bm29

func hasPathSum(root *TreeNode, sum int) bool {
	// write code here
	var flag bool

	var traversal func(root *TreeNode, sum int)
	traversal = func(root *TreeNode, sum int) {
		if flag {
			return
		}
		if root == nil {
			return
		}
		if root.Val == sum && root.Left == nil && root.Right == nil {
			flag = true
			return
		}
		if root.Left != nil {
			traversal(root.Left, sum-root.Val)
		}
		if root.Right != nil {
			traversal(root.Right, sum-root.Val)
		}
	}

	traversal(root, sum)

	return flag
}
