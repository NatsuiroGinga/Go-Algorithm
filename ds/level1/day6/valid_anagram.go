package day6

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	freq := make([]int, 26)
	for _, c := range s {
		freq[c-'a']++
	}
	for _, c := range t {
		freq[c-'a']--
		if freq[c-'a'] == -1 {
			return false
		}
	}
	return true
}
