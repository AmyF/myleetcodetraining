/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	diff := calHeight(root.Left) - calHeight(root.Right)
	return diff >= -1 && diff <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func calHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := calHeight(root.Left)+1, calHeight(root.Right)+1
	if left > right {
		return left
	} else {
		return right
	}
}