package dp

import "math"

func coinChange(coins []int, amount int) int {
	dp := [1e4 + 5]int{0}
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt
	}

	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] = min(dp[j], dp[j-coin]+1)
		}
	}

	if dp[amount] == math.MaxInt || dp[amount] < 0 {
		return -1
	}

	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
