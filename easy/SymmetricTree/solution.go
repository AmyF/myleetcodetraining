/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isSymmetric(root *TreeNode) bool {
    if root == nil { return true }
    return check(root.Left,root.Right)
}

func check(left, right *TreeNode) bool {
    if left == nil || right == nil { return left == right }
    return left.Val == right.Val && check(left.Left,right.Right) && check(left.Right,right.Left)
}