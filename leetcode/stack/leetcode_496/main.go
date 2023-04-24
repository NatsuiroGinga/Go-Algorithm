package main

import "fmt"

func nextGreaterElement(nums1, nums2 []int) []int {
	record := map[int]int{}
	var stack []int
	result := make([]int, len(nums1))

	for i, v := range nums1 {
		record[v] = i
		result[i] = -1
	}

	for _, v := range nums2 {
		for len(stack) > 0 && v > stack[len(stack)-1] {
			beforeValue := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if i, ok := record[beforeValue]; ok {
				result[i] = v
			}
		}
		stack = append(stack, v)
	}

	return result
}

func main() {
	nums1 := []int{2, 4}
	nums2 := []int{1, 2, 3, 4}
	element := nextGreaterElement(nums1, nums2)
	fmt.Println(element)
}
