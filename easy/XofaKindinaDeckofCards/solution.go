func hasGroupsSizeX(deck []int) bool {
	imap := map[int]int{}
	for _, k := range deck {
		imap[k] += 1
	}
	maxLen := 0
	for _, v := range imap {
		if v > maxLen {
			maxLen = v
		}
	}

	for i := 2; i <= maxLen; i++ {
		mapLen := len(imap)
		for _, v := range imap {
			if v%i == 0 {
				mapLen -= 1
			}
		}
		if mapLen == 0 {
			return true
		}
	}
	return false
}