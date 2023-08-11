package main

import (
	"log"
	"math/rand"
)

func main() {
	nums := rand.Perm(10)
	log.Println(nums)
	_quickSort(nums, 0, len(nums)-1)
	log.Println(nums)
}

func _quickSort(nums []int, left, right int) []int {
	if left < right {
		index := partition(nums, left, right)
		_quickSort(nums, left, index-1)
		_quickSort(nums, index+1, right)
	}
	return nums
}

func partition(nums []int, left, right int) int {
	pivot := left
	index := pivot + 1

	for i := index; i <= right; i++ {
		if nums[i] < nums[pivot] {
			nums[i], nums[index] = nums[index], nums[i]
			index++
		}
	}
	nums[pivot], nums[index-1] = nums[index-1], nums[pivot]
	return index - 1
}
