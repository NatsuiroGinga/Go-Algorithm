package day2

func isSubsequence(s string, t string) bool {
	tLen := len(t)
	sLen := len(s)
	j := 0

	for i := 0; i < tLen && j < sLen; i++ {
		if s[j] == t[i] {
			j++
		}
	}

	return j == sLen
}
