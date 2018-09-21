func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxarea := 0
	for i := 0; i < len(height); i++ {
		l, r := height[left], height[right]
		maxarea = max(maxarea, (right-left)*min(l, r))
		if l < r {
			left +=1
		}else{
			right -=1
		}
	}
	return maxarea
}
func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
