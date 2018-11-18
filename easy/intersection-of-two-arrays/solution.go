func intersection(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	tmp := map[int]int{}
	res := []int{}
	for _, n := range nums1 {
		m[n] = 1
	}
	for _, n := range nums2 {
		if m[n] == 1 {
			tmp[n] = 1
		}
	}
	for n, _ := range tmp {
		res = append(res, n)
	}
	return res
}