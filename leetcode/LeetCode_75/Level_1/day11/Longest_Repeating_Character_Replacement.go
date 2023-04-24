package day11

func characterReplacement(s string, k int) int {
	var (
		size     = len(s)
		fre      = make(map[byte]int, 26)
		maxLen   = 0
		maxCount = 0 //  largest count of a single, unique character in the current window
	)

	for left, right := 0, 0; right < size; right++ {
		fre[s[right]]++               // add right char
		if maxCount < fre[s[right]] { // update maxCount
			maxCount = fre[s[right]]
		}

		if right-left+1-maxCount > k { // if window size - maxCount > k, then we need to shrink the window
			fre[s[left]]-- // remove left char
			left++         // shrink window
		}

		if maxLen < right-left+1 { // update maxLen
			maxLen = right - left + 1
		}
	}

	return maxLen
}
