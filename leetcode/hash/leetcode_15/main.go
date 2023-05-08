package leetcode_15

func threeSum(nums []int) [][]int {
	var ans [][]int

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := 0 - nums[i]
		low, high := i+1, len(nums)-1

		for low < high {
			if nums[low]+nums[high] == target {
				ans = append(ans, []int{nums[i], nums[low], nums[high]})

				for low+1 < high && nums[low] == nums[low+1] {
					low++
				}

				for low < high-1 && nums[high] == nums[high-1] {
					high--
				}

				low++
				high--
			} else if nums[low]+nums[high] > target {
				high--
			} else {
				low++
			}
		}
	}

	return ans
}
