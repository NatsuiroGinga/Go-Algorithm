package main

import "fmt"

func largestRectangleArea(heights []int) int {
	heights = append(heights, 0, 0)
	copy(heights[1:], heights)
	heights[0] = 0
	var st []int
	ans := 0

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := range heights {
		for len(st) > 0 && heights[i] < heights[st[len(st)-1]] {
			mid := st[len(st)-1]
			st = st[:len(st)-1]
			if len(st) == 0 {
				break
			}
			left := st[len(st)-1]

			w := i - left - 1
			h := heights[mid]

			ans = max(ans, w*h)
		}
		st = append(st, i)
	}

	return ans
}

func main() {
	a := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea(a))
}
