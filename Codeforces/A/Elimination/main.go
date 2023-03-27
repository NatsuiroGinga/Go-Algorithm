package main

import (
	"fmt"
)

func main() {
	var c, d, n, m, k int
	fmt.Scan(&c, &d, &n, &m, &k)
	limit := n*m - k

	if limit <= 0 {
		fmt.Println(0)
		return
	}

	ans := 0
	if d*n > c { // 正常赛更好
		turn := limit / n               // 正常赛进行几轮
		limit -= turn * n               // 进行完正常赛后剩下的人数
		ans += turn*c + min(limit*d, c) // 正常赛使用题目数 + 剩下的人使用的题目数
	} else { // 加赛更好
		ans += limit * d
	}

	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
