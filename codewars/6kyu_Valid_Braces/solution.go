package kata

func ValidBraces(str string) bool {
	if len(str)%2 == 1 {
		return false
	}
	unko := map[uint8]int{'{': 1, '[': 2, '(': 3, ')': 7, ']': 8, '}': 9}
	tmp := []int{}
	for i := 0; i < len(str); i++ {
		n := unko[str[i]]
		if n < 5 {
			tmp = append(tmp, n)
		} else if len(tmp) <= 0 {
			return false
		} else if tmp[len(tmp)-1]+n == 10 {
			tmp = tmp[:len(tmp)-1]
		} else {
			return false
		}
	}
	return len(tmp) == 0
}
