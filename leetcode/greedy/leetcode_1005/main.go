package leetcode_1005

import (
	"math"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})
	for i := range nums {
		if nums[i] < 0 && k > 0 {
			nums[i] *= -1
			k--
		}
	}
	if k&1 != 0 {
		nums[len(nums)-1] *= -1
	}

	sum := 0
	for _, v := range nums {
		sum += v
	}

	return sum
}
