package day4

import "strings"

func reverseWords(s string) string {
	ans := strings.Builder{}
	words := strings.Split(s, " ")
	flag := false
	for _, word := range words {
		if !flag {
			flag = true
		} else {
			ans.WriteString(" ")
		}

		ans.WriteString(reverse(word))
	}
	return ans.String()
}

func reverse(s string) string {
	var (
		start = 0
		end   = len(s) - 1
		bytes = []byte(s)
	)

	for start < end {
		bytes[start], bytes[end] = bytes[end], bytes[start]
		start++
		end--
	}
	return string(bytes)
}
