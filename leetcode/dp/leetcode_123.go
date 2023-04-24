package dp

func maxProfit3(prices []int) int {
	/*
		dp[i][0]: 不操作
		dp[i][1]: 第一次持有
		dp[i][2]: 第一次不持有
		dp[i][3]: 第二次持有
		dp[i][4]: 第二次不持有
	*/
	dp := [1e5 + 5][4]int{{-prices[0], 0, -prices[0], 0}}
	days := len(prices)

	for i := 1; i < days; i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]-prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]+prices[i])
	}

	return dp[days-1][3]
}
