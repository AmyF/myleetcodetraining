class Solution {
    let unkoT = ["{":1,"[":2,"(":3,"}":9,"]":8,")":7]
    func isValid(_ s: String) -> Bool {
        var tmp: [Int] = []
        for ch in s {
            let t = unkoT[String(ch)]!
            if t < 5 {
                tmp.append(t)
            } else if tmp.count > 0 && tmp.last! + t == 10 {
                tmp.removeLast()
            } else {
                return false
            }
        }      
        return tmp.count == 0
    }
}
