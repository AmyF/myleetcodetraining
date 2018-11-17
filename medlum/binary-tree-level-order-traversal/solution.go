/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		vals := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue[i]
			vals[i] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[size:]
		res = append(res, vals)
	}

	return res
}