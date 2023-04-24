package day8

var ans [][]int
var srcColor int
var dx = []int{-1, 1, 0, 0} // 上下左右
var dy = []int{0, 0, -1, 1}

type Pixel struct {
	x, y int
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	if image[sr][sc] == color {
		return image
	}
	srcColor = image[sr][sc]
	ans = make([][]int, len(image))
	for i := range ans {
		ans[i] = make([]int, len(image[i]))
		copy(ans[i], image[i])
	}
	bfs(image, sr, sc, color)
	//dfs(sr, sc, color, image)
	return ans
}

// dfs
/*func dfs(row, col, color int, image [][]int) {
	if row < 0 || col < 0 || row >= len(ans) || col >= len(ans[0]) || image[row][col] != srcColor || ans[row][col] == color {
		return
	}

	ans[row][col] = color

	dfs(row-1, col, color, image) // 上
	dfs(row+1, col, color, image) // 下
	dfs(row, col-1, color, image) // 左
	dfs(row, col+1, color, image) // 右
}*/

func bfs(image [][]int, sr int, sc int, color int) {
	queue := make([]Pixel, 0)
	queue = append(queue, Pixel{sr, sc})
	ans[sr][sc] = color

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		for i := 0; i < 4; i++ {
			x := p.x + dx[i]
			y := p.y + dy[i]

			if x < 0 || y < 0 || x >= len(ans) || y >= len(ans[0]) || image[x][y] != srcColor || ans[x][y] == color {
				continue
			}

			ans[x][y] = color
			queue = append(queue, Pixel{x, y})
		}
	}
}
