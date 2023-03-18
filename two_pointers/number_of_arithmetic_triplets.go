package two_pointers

// arithmeticTriplets returns the number of arithmetic triplets in nums
/*func arithmeticTriplets(nums []int, diff int) int {
	size := len(nums)
	count := 0 // count the numbers of (i, j, k)

	for i := 0; i < size; i++ { // i is the first element
		for j := i + 1; j < size; j++ { // j is the second element
			if i >= j || nums[j]-nums[i] != diff { // check if i < j and nums[j] - nums[i] == diff
				continue
			}

			// because nums[j] - nums[i] == diff, so nums[j] + diff == nums[k]
			kNum := nums[j] + diff
			// search the kNum in nums[j+1:]
			k := sort.SearchInts(nums[j+1:], kNum)

			// if kNum is found in nums[j+1:], then count++
			if k+j+1 < size && nums[j+1+k] == kNum {
				count++
			}
		}
	}

	return count
}
*/
func arithmeticTriplets(nums []int, diff int) int {
	count := 0
	size := len(nums)
	set := make(map[int]struct{}, size)

	for _, num := range nums {
		set[num] = struct{}{}
	}

	for _, num := range nums {
		_, ok1 := set[num-diff]
		_, ok2 := set[num+diff]
		if ok1 && ok2 {
			count++
		}
	}

	return count
}
