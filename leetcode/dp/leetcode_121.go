package dp

func maxProfit(prices []int) int {
	// dp[i][0]表示第i天持有股票的最大收益, dp[i][1]表示第i天不持有股票的最大收益
	dp := [1e5 + 5][2]int{{-prices[0], 0}}
	days := len(prices)

	for i := 1; i < days; i++ {
		// 第i天 持有股票的最大收益 = max(第i-1天持有股票的最大收益, 第i-1天不持有股票的最大收益-第i天股票价格)
		dp[i][0] = max(dp[i-1][0], -prices[i])
		// 第i天 不持有股票的最大收益 = max(第i-1天不持有股票的最大收益, 第i-1天持有股票的最大收益+第i天股票价格)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}

	return max(dp[days-1][0], dp[days-1][1])
}
