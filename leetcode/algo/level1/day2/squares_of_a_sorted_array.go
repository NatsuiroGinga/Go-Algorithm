package day2

import "math"

func sortedSquares(nums []int) []int {
	var (
		size = len(nums)
		low  = 0
		high = size - 1
		ret  = make([]int, size)
	)

	for i := size - 1; i >= 0; i-- {
		if math.Abs(float64(nums[low])) > math.Abs(float64(nums[high])) {
			ret[i] = nums[low] * nums[low]
			low++
		} else {
			ret[i] = nums[high] * nums[high]
			high--
		}
	}

	return ret
}
