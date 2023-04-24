package dp

func maxProfit2(prices []int) int {
	dp := [3*1e4 + 5][2]int{{-prices[0], 0}}
	days := len(prices)

	for i := 1; i < days; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}

	return max(dp[days-1][0], dp[days-1][1])
}
