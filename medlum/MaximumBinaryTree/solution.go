/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
    if len(nums) == 0 { return nil}
    var root *TreeNode
    for i:= len(nums) - 1; i >= 0; i-- {
        root = tmpFunc(nums[i], root)
    }
    return root
}

func tmpFunc(num int, root *TreeNode) *TreeNode {
    if root == nil {
        return &TreeNode{num,nil,nil}
    }
    if root.Val < num {
        tmp := &TreeNode{num,nil,nil}
        tmp.Right = root
        return tmp
    } else if root.Val > num {
        root.Left = tmpFunc(num, root.Left)
        return root
    }
    return nil
}

