package kata

func LongestConsec(strarr []string, k int) string {
	if len(strarr) == 0 || len(strarr) < k || k <= 0 {
		return ""
	}
	count, flag := 0, 0
	for i := 0; i < len(strarr); i++ {
		tmp := 0
		for j := 0; (i+k) <= len(strarr) && j < k; j++ {
			tmp += len(strarr[i+j])
		}
		if tmp > count {
			flag = i
			count = tmp
		}
	}
	ret := ""
	for i := flag; i < (flag + k); i++ {
		ret += strarr[i]
	}
	return ret
}
