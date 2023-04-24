package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	ans  []string
	path []string
)

func restoreIpAddresses(s string) []string {
	path = make([]string, 0)
	ans = make([]string, 0)
	dfs(0, s)
	return ans
}

func dfs(cur int, s string) {
	if len(path) == 3 {
		str := strings.Join(path, ".")
		last := s[cur:]

		if isValid(last) {
			str += "." + last
		} else {
			return
		}

		ans = append(ans, str)
		return
	}
	for i := cur; i < len(s); i++ {
		// 剪枝, 判断剩余的位数以三个数为单位分割, 能否填充满ip地址
		if (len(s)-(i+1))/3 >= (4 - len(path)) {
			continue
		}

		substr := s[cur : i+1]

		if isValid(substr) {
			path = append(path, substr)
		} else {
			continue
		}

		dfs(i+1, s)

		if sz := len(path); sz > 0 {
			path = path[:sz-1]
		}
	}
}

func isValid(s string) bool {
	if len(s) == 0 || (len(s) > 1 && s[0] == '0') {
		return false
	}
	if num, err := strconv.Atoi(s); err != nil || num > 255 {
		return false
	}
	return true
}

func main() {
	s := "101023"
	addresses := restoreIpAddresses(s)
	fmt.Println(addresses)
}
