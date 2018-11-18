import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return [][]int{}
	}

	tmp := map[string][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l, r := j+1, len(nums)-1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum == target {
					tmp[fmt.Sprint(nums[i], nums[j], nums[l], nums[r])] = []int{nums[i], nums[j], nums[l], nums[r]}
					l += 1
					r -= 1
				} else if sum < target {
					l += 1
				} else {
					r -= 1
				}
			}
		}
	}

	res := [][]int{}
	for _, v := range tmp {
		res = append(res, v)
	}
	return res
}