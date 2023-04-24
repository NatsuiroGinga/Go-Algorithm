package day10

func minCostClimbingStairs(cost []int) int {
	var (
		size = len(cost)
		dp   = make([]int, size+5)
	)

	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < size; i++ {
		dp[i] = MinInt(dp[i-1], dp[i-2]) + cost[i]
	}

	return MinInt(dp[size-1], dp[size-2])
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
