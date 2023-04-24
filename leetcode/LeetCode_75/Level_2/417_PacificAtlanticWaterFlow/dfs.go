package main

import "math"

var dir = [][]int{
	{0, -1},
	{0, 1},
	{-1, 0},
	{1, 0},
}

func pacificAtlantic_2(heights [][]int) [][]int {
	n, m := len(heights), len(heights[0])
	pacific := make([][]bool, n)
	atlantic := make([][]bool, n)
	var ans [][]int

	for i := 0; i < n; i++ {
		pacific[i] = make([]bool, m)
		atlantic[i] = make([]bool, m)
	}

	for i := 0; i < n; i++ {
		dfs(heights, pacific, math.MinInt, i, 0)
		dfs(heights, atlantic, math.MinInt, i, m-1)
	}

	for i := 0; i < m; i++ {
		dfs(heights, pacific, math.MinInt, 0, i)
		dfs(heights, atlantic, math.MinInt, n-1, i)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if pacific[i][j] && atlantic[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}

	return ans
}

func dfs(matrix [][]int, vis [][]bool, height, x, y int) {
	n, m := len(matrix), len(matrix[0])
	if x < 0 || x >= n || y < 0 || y >= m || vis[x][y] || matrix[x][y] < height {
		return
	}
	vis[x][y] = true
	for i := range dir {
		nx := dir[i][0] + x
		ny := dir[i][1] + y
		dfs(matrix, vis, matrix[x][y], nx, ny)
	}
}
