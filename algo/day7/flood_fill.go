package day7

type Node struct {
	x, y int
}

var dx = []int{-1, 1, 0, 0} // 上下左右
var dy = []int{0, 0, -1, 1}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	queue := make([]Node, 0)
	queue = append(queue, Node{sr, sc})
	oldColor := image[sr][sc]
	image[sr][sc] = color

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for i := range dx {
			nx := node.x + dx[i]
			ny := node.y + dy[i]

			if nx < 0 || ny < 0 || nx >= len(image) || ny >= len(image[0]) || image[nx][ny] != oldColor || image[nx][ny] == color {
				continue
			}
			image[nx][ny] = color
			queue = append(queue, Node{nx, ny})
		}
	}

	return image
}
