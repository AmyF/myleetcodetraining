func climbStairs(n int) int {
    switch n {
    case 1: return 1
    case 2: return 2
    case 3: return 3
    default: 
        a, b := 2, 3
        for i:=3; i<n; i++ {
            b, a = b + a, b
        }
        return b
    }
}
