package day1

func pivotIndex(nums []int) (index int) {
	rightSum := 0
	leftSum := 0
	length := len(nums)

	for _, v := range nums {
		rightSum += v
	}

	for i := 0; i < length; i++ {
		rightSum -= nums[i]
		if leftSum == rightSum {
			return i
		}
		leftSum += nums[i]
	}

	return -1
}
