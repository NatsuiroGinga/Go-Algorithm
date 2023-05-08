package leetcode_376

func wiggleMaxLength(nums []int) int {
	up, down := 1, 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			up = down + 1
		} else if nums[i] < nums[i-1] {
			down = up + 1
		}
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	return max(up, down)
}
