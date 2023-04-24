package dp

func numSquares(n int) int {
	var dp [1e4 + 5]int
	var sqr []int
	for i := 1; i <= n; i++ {
		dp[i] = n + 1
	}

	for i := 1; i*i <= n; i++ {
		sqr = append(sqr, i*i)
	}

	for i := 1; i*i <= n; i++ {
		s := i * i
		for j := s; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-s]+1)
		}
	}

	return dp[n]
}
