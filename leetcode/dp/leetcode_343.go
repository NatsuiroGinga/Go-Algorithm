package dp

func integerBreak(n int) int {
	dp := make([]int, n+5)
	dp[0], dp[1] = 0, 0
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			dp[i] = MaxInt(MaxInt(j*(i-j), j*dp[i-j]), dp[i])
		}
	}

	return dp[n]
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
