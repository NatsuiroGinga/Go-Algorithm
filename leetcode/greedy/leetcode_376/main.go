package leetcode_376

func wiggleMaxLength(nums []int) int {
	ans := 0

	for i := 1; i < len(nums)-1; i++ {
		preDif := nums[i] - nums[i-1]
		curDif := nums[i+1] - nums[i]

		if curDif*preDif <= 0 {
			ans++
		}
	}

	ans += 2

	return 0
}
