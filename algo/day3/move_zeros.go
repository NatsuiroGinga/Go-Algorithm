package day3

func moveZeroes(nums []int) {
	zeroCount := 0

	for i, num := range nums {
		if num == 0 {
			zeroCount++
		} else {
			nums[i-zeroCount], nums[i] = nums[i], nums[i-zeroCount]
		}
	}
}
