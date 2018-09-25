class Solution {
    func leastBricks(_ wall: [[Int]]) -> Int {
        var map: [Int:Int] = [:]
        for row in wall {
            var tmp = 0
            for i in 0..<(row.count - 1) {
                tmp += row[i]
                map[tmp] = (map[tmp] ?? 0) + 1
            }
        }
        var count = wall.count
        for (key,val) in map {
            count = min(wall.count - val, count)
        }
        return count 
    }
}
