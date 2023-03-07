package day1

func runningSum(nums []int) []int {
	length := len(nums)
	for i := 1; i < length; i++ {
		nums[i] += nums[i-1]
	}
	return nums
}
