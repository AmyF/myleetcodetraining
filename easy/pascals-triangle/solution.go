func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	ans := [][]int{}
	ans = append(ans, []int{1})
	for i := 1; i < numRows; i++ {
		tmp := make([]int, i+1)
		tmp[0] = 1
		for j := 1; j < i; j++ {
			tmp[j] = ans[i-1][j-1] + ans[i-1][j]
		}
		tmp[i] = 1
		ans = append(ans, tmp)
	}
	return ans
}