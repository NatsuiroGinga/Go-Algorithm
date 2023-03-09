package day6

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	var freS1, freS2 [26]int

	// init freS1
	for _, c := range s1 {
		freS1[c-'a']++
	}

	for left, right := 0, 0; right < len(s2); right++ {
		// init freS2
		freS2[s2[right]-'a']++

		// check if window size is equal to s1
		windowSize := right - left + 1
		if windowSize == len(s1) && isEqual(freS1, freS2) { // if window size is equal to s1 and freS1 == freS2
			return true
		} else if windowSize > len(s1) { // if window size is greater than s1
			freS2[s2[left]-'a']-- // remove left char from freS2
			left++                // move left pointer
		}
	}
	return false
}

// isEqual checks if two arrays are equal
func isEqual(s1 [26]int, s2 [26]int) bool {
	for i := 0; i < 26; i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
