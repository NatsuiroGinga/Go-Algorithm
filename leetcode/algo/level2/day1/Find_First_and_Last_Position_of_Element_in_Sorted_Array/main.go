package Find_First_and_Last_Position_of_Element_in_Sorted_Array

/*
First, we initialize the pointers low and high to point to the first and last elements of the array.
Then, we use the binary search algorithm to find the index of the target element.
If the target element is found, we use two pointers, `l` means left bound, `r` means right bound, to expand the range of the target element.
Finally, we return the range of the target element.
*/
func searchRange(nums []int, target int) []int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			// set left and right pointer to expand the range of target element
			l, r := mid, mid
			// left bound
			for ; l >= 0 && nums[l] == target; l-- {
			}
			// right bound
			for ; r < len(nums) && nums[r] == target; r++ {
			}
			// because the loop will break when l or r is out of range, so we need to add 1 or minus 1 to get the correct index
			return []int{l + 1, r - 1}
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return []int{-1, -1}
}
