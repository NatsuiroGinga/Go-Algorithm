package day6

func canConstruct(ransomNote string, magazine string) bool {
	freqM := make([]int, 26)

	for _, c := range magazine {
		freqM[c-'a']++
	}
	for _, c := range ransomNote {
		freqM[c-'a']--
		if freqM[c-'a'] == -1 {
			return false
		}
	}

	return true
}
