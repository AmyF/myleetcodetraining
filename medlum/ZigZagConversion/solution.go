func convert(s string, numRows int) string {
	if numRows == 1 { return s }
	
	ret := make([]uint8, len(s))
	cycLen := 2*numRows-2
	index := 0
	for i:=0;i<numRows;i++ {
		for j:=i;j<len(s);j+=cycLen {
			ret[index] = s[j]
			index += 1
			if (i != 0 && i != numRows - 1 && j - 2*i + cycLen < len(s)) {
				ret[index] = s[j - 2*i + cycLen]
				index += 1
			}
		}
	}
	return string(ret)
}