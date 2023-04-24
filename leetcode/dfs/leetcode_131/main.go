package main

import "fmt"

var (
	ret  [][]string
	path []string
)

func partition(s string) [][]string {
	ret = make([][]string, 0)
	path = make([]string, 0)
	dfs(0, s)
	return ret
}

func isPalindrome(s string) bool {
	for i, sz := 0, len(s); i < sz/2; i++ {
		if s[i] != s[sz-i-1] {
			return false
		}
	}
	return true
}

func dfs(cur int, s string) {
	if cur == len(s) {
		subset := make([]string, len(path))
		copy(subset, path)
		ret = append(ret, subset)
		return
	}

	for i := cur; i < len(s); i++ {
		substr := s[cur : i+1]

		if isPalindrome(substr) {
			path = append(path, substr)
		} else {
			continue
		}

		dfs(i+1, s)

		if len(path) > 0 {
			path = path[:len(path)-1]
		}
	}
}

func main() {
	s := "aab"
	fmt.Println(partition(s))
}
