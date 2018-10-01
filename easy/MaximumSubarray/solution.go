func maxSubArray(nums []int) int {
    sum := nums[0]
    tmp := 0 
    for i:=0;i<len(nums);i++ {
        if nums[i] + tmp > nums[i] {
            tmp += nums[i]
        } else {
            tmp = nums[i]
        }
        if sum < tmp {
            sum = tmp
        }
    }
    return sum
}
