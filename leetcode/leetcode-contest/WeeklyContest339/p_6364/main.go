package main

func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	dp := make([][]int, len(reward1)+5)
	for i := range dp {
		dp[i] = make([]int, k+5)
	}
	sum2 := 0
	for _, v := range reward2 {
		sum2 += v
	}

	dp[0][0] = sum2

	for i := 1; i <= len(reward1); i++ {
		for j := 1; j <= k; j++ {
			if reward1[i] > reward2[i] {
				dp[i][j] += dp[i-1][j-1] + reward1[i-1]
			} else {
				dp[i][j] += dp[i-1][j] + reward2[i-1]
			}
		}
	}
	return dp[len(reward2)][k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
