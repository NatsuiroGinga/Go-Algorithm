package leetcode_695

func maxAreaOfIsland(grid [][]int) int {
	ans := 0
	rows := len(grid)
	cols := len(grid[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if area := dfs(i, j, grid); area > ans {
				ans = area
			}
		}
	}

	return ans
}

func dfs(x, y int, grid [][]int) int {
	if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == 1 {
		grid[x][y] = 0
		return 1 + dfs(x-1, y, grid) + dfs(x+1, y, grid) + dfs(x, y-1, grid) + dfs(x, y+1, grid)
	}
	return 0
}
