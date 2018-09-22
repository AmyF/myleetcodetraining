func sortArrayByParity(A []int) []int {
    B := make([]int, len(A))
    left, right := 0, len(A)-1
    for i := 0; i < len(A); i++ {
        if right == left {
            B[right] = A[i]
            break
        }
        if A[i] & 1 == 1 {
            B[right] = A[i]
            right -= 1
        } else {
            B[left] = A[i]
            left += 1
        }
    }
    return B
}
