import "sort"

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}

	res := nums[0] + nums[1] + nums[2]
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			tmp := nums[i] + nums[l] + nums[r]
			if abs(target-tmp) < abs(target-res) {
				res = tmp
			}
			if tmp < target {
				l += 1
			} else {
				r -= 1
			}
		}
	}
	return res
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
} 