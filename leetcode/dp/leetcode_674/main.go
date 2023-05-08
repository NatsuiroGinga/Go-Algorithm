package leetcode_674

func findLengthOfLCIS(nums []int) int {
	cur, ans := 1, 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			cur++
			if cur > ans {
				ans = cur
			}
		} else {
			cur = 1
		}
	}

	return ans
}
