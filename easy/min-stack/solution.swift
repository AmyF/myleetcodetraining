
class MinStack {
    private struct Data {
        let data: Int
        let min: Int
    }
    private var datas: [Data]
    
    /** initialize your data structure here. */
    init() {
        datas = []
    }
    
    func push(_ x: Int) {
        if datas.isEmpty {
            datas.append(.init(data: x, min: x))
        } else {
            let min = getMin()
            if x < min {
                datas.append(.init(data: x, min: x))
            } else {
                datas.append(.init(data: x, min: min))
            }
        }
    }
    
    func pop() {
        datas.popLast()
    }
    
    func top() -> Int {
        return datas.last!.data
    }
    
    func getMin() -> Int {
        return datas.last!.min
    }
}

/**
 * Your MinStack object will be instantiated and called as such:
 * let obj = MinStack()
 * obj.push(x)
 * obj.pop()
 * let ret_3: Int = obj.top()
 * let ret_4: Int = obj.getMin()
 */