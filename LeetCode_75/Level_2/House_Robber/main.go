package main

func rob(nums []int) int {
	pre := nums[0]
	pos := 0

	for i := 1; i < len(nums); i++ {
		x := Max(pos, pre+nums[i])
		pre = pos
		pos = x
	}

	return pos
}

/*func recur(nums []int, i int) int {
	if i < 0 {
		return 0
	}

	if dp[i] >= 0 {
		return dp[i]
	}

	max := Max(recur(nums, i-2)+nums[i], recur(nums, i-1))
	dp[i] = max

	return max
}*/

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
