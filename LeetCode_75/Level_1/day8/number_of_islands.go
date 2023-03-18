package day8

var dir = [][]int{
	{-1, 0}, // up
	{1, 0},  // down
	{0, -1}, // left
	{0, 1},  // right
}

func numIslands(grid [][]byte) int {
	ret := 0
	for i, bytes := range grid {
		for j, b := range bytes {
			if b == '1' {
				ret++
				dfs(grid, i, j)
			}
		}
	}

	return ret
}

func dfs(grid [][]byte, row, col int) {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) || grid[row][col] == '0' || grid[row][col] == 'x' {
		return
	}
	grid[row][col] = 'x'

	for _, ints := range dir {
		dfs(grid, row+ints[0], col+ints[1])
	}
}
