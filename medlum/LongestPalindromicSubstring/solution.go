func longestPalindrome(s string) string {
	length := len(s)
	if length == 0 {
		return ""
	}
	ret := s[:1]
	for i := 0; i < length; i++ {
		// 单数
		for j := 1; i-j >= 0 && i+j < length; j++ {
			if s[i-j] == s[i+j] {
				if len(ret) < 2*j+1 {
					ret = s[i-j : i+j+1]
				}
			} else {
				break
			}
		}
		// 双数
		for j := 1; i-(j-1) >= 0 && i+j < length; j++ {
			if s[i-(j-1)] == s[i+j] {
				if len(ret) < 2*j {
					ret = s[i-(j-1) : i+j+1]
				}
			} else {
				break
			}
		}
	}
	return ret
}