package main

import (
	"bytes"
	"fmt"
)

var c [3][101]bool
var ans [][]string

func solveNQueens(n int) [][]string {
	c = [3][101]bool{}
	ans = make([][]string, 0)
	chessboard := make([]string, n)
	dfs(chessboard, 0, n) // 第0行开始
	return ans
}

func dfs(chessboard []string, cur, n int) {
	if cur == n {
		var solution []string
		for _, v := range chessboard {
			solution = append(solution, v)
		}
		ans = append(ans, solution)
		return
	}

	line := bytes.Repeat([]byte("."), n)

	for i := 0; i < n; i++ {
		if c[0][i] || c[1][cur+i] || c[2][n+cur-i] {
			continue
		}

		c[0][i], c[1][cur+i], c[2][n+cur-i] = true, true, true
		line[i] = 'Q'

		dfs(chessboard, cur+1, n)

		line[i] = '.'
		c[0][i], c[1][cur+i], c[2][n+cur-i] = false, false, false
	}
}

func main() {
	n := 4
	queens := solveNQueens(n)
	fmt.Println(queens)
}
