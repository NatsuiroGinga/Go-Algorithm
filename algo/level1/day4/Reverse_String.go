package day4

func reverseString(s []byte) {
	var (
		size  = len(s)
		start = 0
		end   = size - 1
	)

	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}
