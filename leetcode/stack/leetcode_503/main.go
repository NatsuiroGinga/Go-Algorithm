package main

import "fmt"

func nextGreaterElements(nums []int) []int {
	result := make([]int, len(nums))
	var stack []int
	n := len(nums)

	for i := range result {
		result[i] = -1
	}

	for i, sz := 0, 2*n; i < sz; i++ {
		for len(stack) > 0 && nums[i%n] > nums[stack[len(stack)-1]] {
			index := stack[len(stack)-1]
			result[index] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%n)
	}

	return result
}

func main() {
	nums := []int{1, 2, 3, 4, 3}
	elements := nextGreaterElements(nums)
	fmt.Println(elements)
}
