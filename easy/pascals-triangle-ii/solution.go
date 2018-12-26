func getRow(rowIndex int) []int {
	ans := make([]int, rowIndex+1)
	num, j := 1, 1
	for i := 1; i < rowIndex; i++ {
		num *= rowIndex - i + 1
		num /= j
		ans[i] = num
		j += 1
	}
	ans[0], ans[rowIndex] = 1, 1
	return ans
}