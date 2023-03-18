package main

import (
	"unicode"
)

var tmp []rune
var ans []string

func letterCasePermutation(s string) []string {
	tmp = make([]rune, 0)
	ans = make([]string, 0)
	tmp = []rune(s)

	per(len(s), 0)
	return ans
}

func per(n, cur int) {
	if cur == n {
		ans = append(ans, string(tmp))
		return
	}

	for i := cur; i < n; i++ {
		tmp[i] = trans(tmp[i])
		per(n, cur+1)
		tmp[i] = trans(tmp[i])
	}
}

func trans(char rune) rune {
	if unicode.IsDigit(char) {
		return char
	} else if unicode.IsLower(char) {
		return unicode.ToUpper(char)
	}
	return unicode.ToLower(char)
}
