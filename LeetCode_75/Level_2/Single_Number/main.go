package main

func singleNumber(nums []int) int {
	var xor int
	for _, num := range nums {
		xor ^= num
	}
	return xor
}
