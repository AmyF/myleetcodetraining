func generateParenthesis(n int) []string {
	ans := []string{}
	if n == 0 {
		ans = append(ans, "")
	} else {
		for i:=0;i<n;i++ {
			left := generateParenthesis(i)
			for j:=0;j<len(left);j++ {
				right := generateParenthesis(n-1-i)
				for k:=0;k<len(right);k++ {
					ans = append(ans, "("+left[j]+")"+right[k])
				}
			}
		}
	}
	return ans
}
