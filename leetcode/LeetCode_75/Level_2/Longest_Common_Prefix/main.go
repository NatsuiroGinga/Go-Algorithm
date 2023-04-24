package main

import (
	"sort"
)

func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}
	sort.Strings(strs)
	first := strs[0]
	last := strs[len(strs)-1]

	var c int
	for c = 0; c < len(first) && first[c] == last[c]; c++ {
	}

	if c == 0 {
		return ""
	}

	return first[:c]
}
