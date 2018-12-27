func maxProfit(prices []int) int {
	ans := 0
	for i := 1; i < len(prices); i++ {
		if prices[i-1] < prices[i] {
			ans += prices[i] - prices[i-1]
		}
	}
	return ans
}