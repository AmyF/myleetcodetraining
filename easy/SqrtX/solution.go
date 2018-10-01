func mySqrt(x int) int {
    if x == 0 { return 0 }
    mid := x
    for {
        if x / mid < mid {
            mid = (mid + x/mid)/2
        } else {
            break
        }
    }
    return mid
}
