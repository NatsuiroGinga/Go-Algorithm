package day7

func maxAreaOfIsland(grid [][]int) int {
	var maxArea = 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] == 1 {
				maxArea = MaxInt(maxArea, dfs(row, col, grid))
			}
		}
	}
	return maxArea
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(row, col int, grid [][]int) int {
	if row >= 0 && row < len(grid) && col >= 0 && col < len(grid[0]) && grid[row][col] == 1 {
		grid[row][col] = 0
		return 1 + dfs(row-1, col, grid) + dfs(row+1, col, grid) + dfs(row, col-1, grid) + dfs(row, col+1, grid)
	}
	return 0
}
