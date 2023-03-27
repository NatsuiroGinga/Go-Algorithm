package main

func minOperations(nums []int, queries []int) []int64 {
	var ans []int64
	var numSum int64

	for _, num := range nums {
		numSum += int64(num)
	}

	for _, query := range queries {
		l := int64(len(nums))
		ans = append(ans, numSum-int64(query)*l)
	}

	return ans
}
