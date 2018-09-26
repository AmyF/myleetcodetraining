func isPalindrome(s string) bool {
    length := len(s)
	if length == 0 { return true }
	check := func(ch uint8) uint8 {
		if (ch >= '0' && ch <= '9') {
			return ch
		}
		if ch >= 'A' && ch <= 'Z' {
			return ch + ('a'-'A')
		}
		if ch >= 'a' && ch <= 'z' {
			return ch
		}
		return 0
	}
	for i,j:=0,length-1;i<j; {
		var a uint8 = 0
		for {
			if i > j {
				break
			}
			a = check(s[i])
			if a != 0 {
				break
			}
			i++
		}
		var b uint8 = 0
		for {
			if j < i {
				break
			}
			b = check(s[j])
			if b != 0 {
				break
			}
			j--
		}
		if i > j && a != 0 && b != 0 {
			return false
		}
		if a != b {
			return false
		}
		i++
		j--
	}
	return true
}
