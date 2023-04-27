package day1

import "math"

func containsDuplicate(nums []int) bool {
	flag := make(map[int]bool, math.MaxInt)

	for _, v := range nums {
		if !flag[v] {
			flag[v] = true
		} else {
			return true
		}
	}

	return false
}
