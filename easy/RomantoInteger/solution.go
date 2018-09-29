func romanToInt(s string) int {
    num := 0
    for i:=0;i<len(s);i++ {
        ch := s[i]
        switch ch {
            case 'I': 
            num += 1
            case 'V':
            num += 5
            if i > 0 && s[i-1] == 'I' {
                num -= 2
            }
            case 'X':
            num += 10
            if i > 0 && s[i-1] == 'I' {
                num -= 2
            }
            case 'L':
            num += 50
            if i > 0 && s[i-1] == 'X' {
                num -= 20
            }
            case 'C':
            num += 100
            if i > 0 && s[i-1] == 'X' {
                num -= 20
            }
            case 'D':
            num += 500
            if i > 0 && s[i-1] == 'C' {
                num -= 200
            }
            case 'M':
            num += 1000
            if i > 0 && s[i-1] == 'C' {
                num -= 200
            }
        }
    }
    return num
}
