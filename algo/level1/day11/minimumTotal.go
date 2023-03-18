package day11

func minimumTotal(triangle [][]int) int {
	var (
		n  = len(triangle)
		dp = make([]int, n+5)
	)
	for i := 0; i < n; i++ {
		dp[i] = triangle[n-1][i]
	}

	for i := n - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}

	return dp[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
