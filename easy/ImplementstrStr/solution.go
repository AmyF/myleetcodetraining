func strStr(haystack string, needle string) int {
    if needle == "" { return 0 }
    al, bl := len(haystack), len(needle)
    
    for i:=0;i<al;i++ {
        if al - i < bl {
			return -1
		}
        for j:=0;j<bl;j++ {
            if haystack[i+j] != needle[j] {
                break
            }
            if j == bl - 1 {
                return i
            }
        }
    }
    return -1
}
