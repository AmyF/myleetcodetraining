func removeDuplicates(nums []int) int {
    length := len(nums)
    if length == 0 { return 0 }
    count := 1
    for i := 1; i < length; i++ {
        if nums[i] != nums[i-1] {
            nums[count] = nums[i]
            count += 1
        }
    }
    return count
}
