package day1

func searchInsert(nums []int, target int) int {
	var (
		size = len(nums)
		low  = 0
		high = size - 1
	)

	if size == 0 {
		return 0
	}

	for low < high {
		mid := low + (high-low)/2
		if target > nums[mid] {
			low = mid + 1
		} else if target < nums[mid] {
			high = mid
		} else {
			return mid
		}
	}

	if target <= nums[low] {
		return low
	}

	return low + 1
}
