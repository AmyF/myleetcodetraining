class Solution {
    func maxIncreaseKeepingSkyline(_ grid: [[Int]]) -> Int {
        var tb: [Int] = [Int](repeating: 0, count: grid.first?.count ?? 0)
        var lr: [Int] = [Int](repeating: 0, count: grid.count)
        for (i, line) in grid.enumerated() {
            lr[i] = line.max() ?? 0
            for (j, height) in line.enumerated() {
                tb[j] = max(tb[j], height)
            }
        }
        
        var count = 0
        for (i, line) in grid.enumerated() {
            for (j, height) in line.enumerated() {
                count += min(tb[i], lr[j]) - height
            }
        }
        return count
    }
}
