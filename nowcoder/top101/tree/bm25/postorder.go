package bm25

func postorderTraversal(root *TreeNode) []int {
	// write code here
	var inorder func(root *TreeNode) []int
	inorder = func(root *TreeNode) []int {
		if root == nil {
			return nil
		}
		l := inorder(root.Left)
		l = append(l, inorder(root.Right)...)
		l = append(l, root.Val)
		return l
	}

	return inorder(root)
}
