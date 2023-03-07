package day5

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	var (
		maxCur   = 0
		maxSoFar = 0
		size     = len(prices)
	)
	for i := 1; i < size; i++ {
		maxCur += prices[i] - prices[i-1]
		maxCur = max(0, maxCur)
		maxSoFar = max(maxSoFar, maxCur)
	}

	return maxSoFar
}
