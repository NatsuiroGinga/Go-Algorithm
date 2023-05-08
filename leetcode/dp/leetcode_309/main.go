package main

func maxProfit(prices []int) int {
	dp := [5000][4]int{{-prices[0], 0, 0, 0}}

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], max(dp[i-1][3]-prices[i], dp[i-1][1]-prices[i]))
		dp[i][1] = max(dp[i-1][1], dp[i-1][3])
		dp[i][2] = dp[i-1][0] + prices[i]
		dp[i][3] = dp[i-1][2]
	}

	return max(max(dp[len(prices)-1][1], dp[len(prices)-1][2]), dp[len(prices)-1][3])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
