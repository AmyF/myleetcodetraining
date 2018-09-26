func reverseWords(s string) string {
	var buf bytes.Buffer
	sarr := strings.Split(s," ")
	for i:=0;i<len(sarr);i++ {
		buf.WriteString(reverseWord(sarr[i]))
		if i < len(sarr)-1 {
			buf.WriteString(" ")
		}
	}

	return buf.String()
}

func reverseWord(s string) string {
	var buf bytes.Buffer
	l := len(s)
	for i := l-1; i >= 0; i-- {
		buf.WriteRune(rune(s[i]))
	}
	return buf.String()
}
