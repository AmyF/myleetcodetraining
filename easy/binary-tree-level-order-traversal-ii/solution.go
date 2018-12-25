/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	fakeAns := [][]int{}
	if root == nil {
		return fakeAns
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		tmp := []int{}
		for i, n := 0, len(q); i < n; i++ {
			node := q[0]
			q = q[1:]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		fakeAns = append(fakeAns, tmp)
	}

	ans := [][]int{}
	for i := len(fakeAns) - 1; i >= 0; i = i - 1 {
		ans = append(ans, fakeAns[i])
	}

	return ans
}