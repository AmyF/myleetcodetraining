func addBinary(a string, b string) string {
    carry := false
    lenOfA := len(a)
    lenOfB := len(b)
    if lenOfA < lenOfB {
        tmp := a
        a = b
        b = tmp
        lenOfA = len(a)
        lenOfB = len(b)
    }
    
    s := make([]byte, 0, lenOfA + 1)
    for i,j := lenOfA-1,lenOfB-1; i >= 0 && j >= 0; i,j=i-1,j-1 {
        if a[i] == '1' && b[j] == '1' {
            if carry {
                s = append(s, '1')
            } else {
                carry = true
                s = append(s, '0')
            }
        } else if (a[i]=='1' && b[j]=='0') || (a[i]=='0'&&b[j]=='1') {
            if carry {
                s = append(s, '0')
            } else {
                s = append(s, '1')
            }
        } else {
            if carry {
                carry = false
                s = append(s, '1')
            } else {
                s = append(s, '0')
            }
        }
    }
    
    diff := lenOfA - lenOfB
    if diff == 0 {
        if carry {
            s = append(s, '1')
        }
    } else {
        for i := diff-1; i >= 0; i-- {
            if a[i] == '1' && carry {
                s = append(s, '0')
            } else if carry {
                s = append(s, '1')
                carry = false
            } else {
                s = append(s, a[i])
            }
        }
        if carry {
            s = append(s, '1')
        }
    }
    
    for i := 0; i < len(s)/2; i++ {
        s[i],s[len(s)-i-1] = s[len(s)-i-1],s[i]
    }
    
    return string(s)
}
