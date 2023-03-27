package main

import (
	"bufio"
	"fmt"
	"os"
)

var n, m, k int
var maze [][]byte
var dir = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func main() {
	gin := bufio.NewReader(os.Stdin)
	gout := bufio.NewWriter(os.Stdout)
	defer gout.Flush()

	fmt.Fscan(gin, &n, &m, &k)
	maze = make([][]byte, n+5)

	pointCount := 0

	for i := 1; i <= n; i++ {
		var line string
		fmt.Fscan(gin, &line)
		maze[i] = make([]byte, m+5)

		for j := 1; j <= m; j++ {
			if line[j-1] == '.' {
				maze[i][j] = 'X'
				pointCount++
			} else {
				maze[i][j] = line[j-1]
			}
		}
	}

	k = pointCount - k

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if maze[i][j] == 'X' {
				dfs(i, j)

				for l := 1; l <= n; l++ {
					fmt.Fprintln(gout, string(maze[l][1:m+1]))
				}

				return
			}
		}
	}
}

func dfs(row, col int) {
	if row > n || row < 1 || col < 1 || col > m || k <= 0 || maze[row][col] != 'X' {
		return
	}

	k--
	maze[row][col] = '.'

	for i := 0; i < 4; i++ {
		nx := dir[i][0] + row
		ny := dir[i][1] + col
		dfs(nx, ny)
	}
}
