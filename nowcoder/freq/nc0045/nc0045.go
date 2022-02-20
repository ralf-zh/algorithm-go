package nc0045

/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 *
 * @param root TreeNode类 the root of binary tree
 * @return int整型二维数组
 */
func threeOrders(root *TreeNode) [][]int {
	// write code here
	return [][]int{preOrder(root), inOrder(root), postOrder(root)}
}

func preOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	r := append([]int{root.Val}, preOrder(root.Left)...)
	r = append(r, preOrder(root.Right)...)
	return r
}

func inOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	r := append(inOrder(root.Left), root.Val)
	r = append(r, inOrder(root.Right)...)
	return r
}

func postOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	r := append(postOrder(root.Left), postOrder(root.Right)...)
	r = append(r, root.Val)
	return r
}
