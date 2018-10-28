func binaryGap(N int) int {
	flag, ans := -1, 0
	for i := 0; i < 32; i++ {
		if N&1 > 0 {
			if flag >= 0 {
				if i-flag > ans {
					ans = i - flag
				}
			}
			flag = i
		}
		N = N >> 1
	}
	return ans
}