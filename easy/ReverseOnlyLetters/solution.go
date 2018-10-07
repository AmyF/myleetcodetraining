func reverseOnlyLetters(S string) string {
	r := make([]uint8, len(S))
	left, right := 0, len(S)-1
	isW := func(c uint8) bool {
		return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')
	}
	for left <= right {
		n := S[left]
		if isW(n) {
			m := S[right]
			for !isW(m) {
				r[right] = m
				right -= 1
				m = S[right]
			}
			r[right] = n
			r[left] = m
			right -= 1
		} else {
			r[left] = n
		}
		left += 1
	}
	return string(r)
}