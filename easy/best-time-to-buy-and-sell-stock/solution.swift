class Solution {
    func maxProfit(_ prices: [Int]) -> Int {
        var min = Int.max
        var max = 0
        for p in prices {
            if p < min {
                min = p
            }
            else if p - min > max {
                max = p - min
            }
        }
        return max
    }
}