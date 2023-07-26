package main

import "fmt"

func main() {
	nums := []int{1, 9, 4, 5, 0}
	fmt.Println(sortArray(nums))
}

func sortArray(nums []int) []int {
	return quicksort(nums)
}

func quicksort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	pivot := nums[0]
	var left, right []int

	for _, num := range nums[1:] {
		if num < pivot {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}

	return append(quicksort(left),
		append([]int{pivot}, quicksort(right)...)...)
}
