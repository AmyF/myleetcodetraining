func isMonotonic(A []int) bool {
	if len(A) <= 1 {
		return true
	}
	flag := 0
	for i := 1; i < len(A); i++ {
		tmp := 0
		if A[i] > A[i-1] {
			tmp = 1
		} else if A[i] < A[i-1] {
			tmp = -1
		}

		if tmp != 0 {
			if tmp != flag && flag != 0 {
				return false
			}
			flag = tmp
		}
	}
	return true
}