func plusOne(digits []int) []int {
    carry := true
    for i:=len(digits)-1;i>=0;i-- {
        if carry {
            digits[i] += 1
            carry = false
        } else {
            break
        }
        if digits[i] > 9 {
            digits[i] = 0
            carry = true
        }
    }
    if carry {
        return append([]int{1}, digits...)
    } else {
        return digits
    }
}
