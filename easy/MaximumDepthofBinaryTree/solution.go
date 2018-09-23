/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    leftNum := 1
    if root.Left != nil {
        leftNum = leftNum + maxDepth(root.Left)
    }
    rightNum := 1
    if root.Right != nil {
        rightNum = rightNum + maxDepth(root.Right)
    }

    if leftNum > rightNum {
        return leftNum
    } else {
        return rightNum
    }
}
