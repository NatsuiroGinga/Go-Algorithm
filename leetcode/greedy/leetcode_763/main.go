package main

import "fmt"

func main() {
	s := "eccbbbbdec"
	fmt.Println(partitionLabels(s))
}

func partitionLabels(s string) []int {
	var dis [26]int
	for i := 0; i < len(s); i++ {
		dis[s[i]-'a'] = i
	}

	ans := make([]int, 0)
	bound := 0
	sz := 0
	for i := 0; i < len(s); i++ {
		sz++
		if dis[s[i]-'a'] > bound {
			bound = dis[s[i]-'a']
		}
		if i == bound {
			ans = append(ans, sz)
			sz = 0
		}
	}

	return ans
}
