func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	tmp := map[string][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		// 剪枝优化
		if i >= 1 {
			for nums[i] == nums[i-1] && i < len(nums)-2 {
				i++
				continue
			}
		}

		l, r := i+1, len(nums)-1
		for l < r {
			if nums[l]+nums[r] > -nums[i] {
				r -= 1
				continue
			} else if nums[l]+nums[r] < -nums[i] {
				l += 1
				continue
			} else {
				tmp[fmt.Sprint(nums[i], nums[l], nums[r])] = []int{nums[i], nums[l], nums[r]}
				l, r = l+1, r-1
			}
		}
	}
	ret := [][]int{}
	for _, v := range tmp {
		ret = append(ret, v)
	}
	return ret
}