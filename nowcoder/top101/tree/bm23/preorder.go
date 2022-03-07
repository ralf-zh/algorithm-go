package bm23

func preorderTraversal(root *TreeNode) []int {
	// write code here
	var preorder func(root *TreeNode) []int
	preorder = func(root *TreeNode) []int {
		if root == nil {
			return nil
		}
		l := append([]int{root.Val}, preorder(root.Left)...)
		l = append(l, preorder(root.Right)...)
		return l
	}

	return preorder(root)
}
