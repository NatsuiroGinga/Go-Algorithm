package day2

func rotate(nums []int, k int) {
	size := len(nums)
	reverse(nums, 0, size-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, size-1)
}

func reverse(nums []int, start int, end int) {
	for start < end {
		tmp := nums[start]
		nums[start] = nums[end]
		nums[end] = tmp
		start++
		end--
	}
}
