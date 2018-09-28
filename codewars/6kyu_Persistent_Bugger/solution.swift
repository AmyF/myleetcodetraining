func persistence(for num: Int) -> Int {
    if num < 10 { return 0 }
    var tmp = num, result = 1
    while tmp > 0 {
        result *= tmp % 10
        tmp /= 10
    }
    return 1 + persistence(for:result)
}
