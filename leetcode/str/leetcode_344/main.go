package leetcode_344

func reverseString(s []byte) {
	for low, high := 0, len(s)-1; low < high; low, high = low+1, high-1 {
		s[low], s[high] = s[high], s[low]
	}
}
