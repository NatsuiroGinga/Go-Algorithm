package main

import (
	"fmt"
)

func maxSlidingWindow(nums []int, k int) []int {
	var ans, window []int

	for _, num := range nums {
		window = append(window, num)

		if len(window) == k {
			max := window[0]
			ans = append(ans, max)
			window = window[1:]
		}
	}

	return ans
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))
}
