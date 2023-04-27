package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	var ans [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			low, high := j+1, len(nums)-1

			for low < high {
				sum := nums[i] + nums[j] + nums[low] + nums[high]
				if sum > target {
					high--
				} else if sum < target {
					low++
				} else {
					ans = append(ans, []int{nums[i], nums[j], nums[low], nums[high]})

					for low+1 < high && nums[low] == nums[low+1] {
						low++
					}

					for low < high-1 && nums[high] == nums[high-1] {
						high--
					}

					low++
					high--
				}
			}
		}
	}

	return ans
}

func main() {
	nums := []int{-2, -1, -1, 1, 1, 2, 2}
	target := 0
	sum := fourSum(nums, target)
	fmt.Println(sum)
}
