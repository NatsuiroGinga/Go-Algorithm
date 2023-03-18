package day9

type Node struct {
	x, y int
}

var dir = [][]int{
	{0, 1},  // right
	{0, -1}, // left
	{1, 0},  // down
	{-1, 0}, // up
}

func updateMatrix(mat [][]int) [][]int {
	queue := make([]Node, 0)
	rows := len(mat)
	cols := len(mat[0])

	for i, line := range mat {
		for j, v := range line {
			if v == 0 {
				queue = append(queue, Node{i, j})
			} else {
				mat[i][j] = -1
			}
		}
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, line := range dir {
			nx := line[0] + node.x
			ny := line[1] + node.y

			if nx >= 0 && nx < rows && ny >= 0 && ny < cols && mat[nx][ny] == -1 {
				mat[nx][ny] = mat[node.x][node.y] + 1
				queue = append(queue, Node{nx, ny})
			}
		}
	}

	return mat
}
