package main

func main() {
}

func findBall(grid [][]int) []int {
	var (
		rows = len(grid)
		cols = len(grid[0])
		ans  = make([]int, 0, cols)
	)

	for i := 0; i < cols; i++ {
		var next int
		cur := i

		for j := 0; j < rows; j++ {
			next = cur + grid[j][cur]
			if next < 0 || next >= cols || grid[j][cur] != grid[j][next] {
				cur = -1
				break
			}
			cur = next
		}

		ans = append(ans, cur)
	}

	return ans
}
