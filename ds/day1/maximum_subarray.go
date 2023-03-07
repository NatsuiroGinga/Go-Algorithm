package day1

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	size := len(nums)
	dp := make([]int, size)
	dp[0] = nums[0]

	for i := 1; i < size; i++ {
		tmp := 0
		if dp[i-1] > 0 {
			tmp = dp[i-1]
		}
		dp[i] = tmp + nums[i]

		if maxSum < dp[i] {
			maxSum = dp[i]
		}
	}

	return maxSum
}
