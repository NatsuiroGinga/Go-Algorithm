package day10

func uniquePaths(m int, n int) int {
	var dp = make([][]int, m+5)
	for i := range dp {
		dp[i] = make([]int, n+5)
	}

	dp[1][1] = 1

	for i := 1; i <= m; i++ {
		for j := 1; i <= n; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}

	return dp[m][n]
}
