package dp

func maxProfit4(k int, prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2*k+1)
	}

	for j := 1; j < 2*k; j += 2 {
		dp[0][j] = -prices[0]
	}

	for i := 1; i < len(prices); i++ {
		for j := 0; j < 2*k; j += 2 {
			dp[i][j+1] = max(dp[i-1][j+1], dp[i-1][j]-prices[i])   // 持有
			dp[i][j+2] = max(dp[i-1][j+2], dp[i-1][j+1]+prices[i]) // 不持有
		}
	}

	return dp[len(prices)-1][2*k]
}
