package day1

func search(nums []int, target int) int {
	var (
		size  = len(nums)
		left  = 0
		right = size - 1
	)

	for left < right {
		mid := left + (right-left+1)/2
		if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
