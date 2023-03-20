package main

/*
1. Binary Search
2. If nums[mid] > nums[high], the minimum must be in the right half
3. If nums[mid] < nums[high], the minimum must be in the left half
4. If nums[mid] == nums[high], the minimum must be in the left half
5. The loop will end when low == high
6. Return nums[low]
*/
func findMin(nums []int) int {
	n := len(nums)
	low, high := 0, n-1
	for low < high {
		mid := (low + high) / 2
		if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return nums[low]
}

func main() {

}
