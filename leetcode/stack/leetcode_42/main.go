package main

import (
	"fmt"
)

func trap(height []int) int {
	var stack []int
	ans := 0

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := range height {
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			low := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if len(stack) > 0 {
				left := stack[len(stack)-1]
				ans += (min(height[left], height[i]) - height[low]) * (i - left - 1)
			}
		}

		if len(stack) > 0 && height[i] == height[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	return ans
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	ans := trap(height)
	fmt.Println(ans)
}
