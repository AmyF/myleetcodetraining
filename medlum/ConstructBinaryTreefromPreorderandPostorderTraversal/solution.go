/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructFromPrePost(pre []int, post []int) *TreeNode {
	N := len(pre)
	if N == 0 {
		return nil
	}

	root := &TreeNode{Val: pre[0]}
	if N == 1 {
		return root
	}

	L := 0
	for i := 0; i < N; i++ {
		if post[i] == pre[1] {
			L = i + 1
		}
	}

	root.Left = constructFromPrePost(pre[1:L+1], post[0:L])
	root.Right = constructFromPrePost(pre[L+1:N], post[L:N-1])

	return root
}