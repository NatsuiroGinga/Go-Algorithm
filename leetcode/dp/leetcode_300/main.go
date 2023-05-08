package leetcode_300

import (
	"sort"
)

func lengthOfLIS(nums []int) int {
	sub := []int{nums[0]}
	for _, num := range nums {
		if num > sub[len(sub)-1] {
			sub = append(sub, num)
		} else {
			idx := sort.Search(len(sub), func(i int) bool {
				return sub[i] >= num
			})
			sub[idx] = num
		}
	}
	return len(sub)
}
