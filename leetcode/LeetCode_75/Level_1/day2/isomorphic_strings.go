package day2

func isIsomorphic(s string, t string) bool {
	var m1, m2 [256]int
	length := len(s)

	for i := 0; i < length; i++ {
		if m1[s[i]] != m2[t[i]] {
			return false
		}
		m1[s[i]] = i + 1
		m2[t[i]] = i + 1
	}

	return false
}
