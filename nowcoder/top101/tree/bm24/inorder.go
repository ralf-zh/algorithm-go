package bm24

func inorderTraversal(root *TreeNode) []int {
	// write code here
	var inorder func(root *TreeNode) []int
	inorder = func(root *TreeNode) []int {
		if root == nil {
			return nil
		}
		l := inorder(root.Left)
		l = append(l, root.Val)
		l = append(l, inorder(root.Right)...)
		return l
	}

	return inorder(root)
}
