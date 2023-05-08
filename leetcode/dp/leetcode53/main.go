package leetcode53

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	pos := nums[0]
	pre := 0

	for i, sz := 1, len(nums); i < sz; i++ {
		pos = max(pre, 0) + nums[i]
		maxSum = max(maxSum, pos)
		pre = pos
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
