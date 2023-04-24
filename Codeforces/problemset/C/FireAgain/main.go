package main

import (
	"bufio"
	"fmt"
	"os"
)

var n, m, k int

var dir = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

type Point struct {
	X, Y int
}

func main() {
	input, _ := os.Open("input.txt")
	output, _ := os.Create("output.txt")
	ansX, ansY := 0, 0
	gin := bufio.NewReader(input)
	gout := bufio.NewWriter(output)
	defer input.Close()
	defer output.Close()
	defer gout.Flush()

	fmt.Fscan(input, &n, &m, &k)
	queue := make([]Point, 0)
	var vis [2005][2005]byte

	for i := 0; i < k; i++ {
		point := Point{}
		fmt.Fscan(gin, &point.X, &point.Y)
		queue = append(queue, point)
		vis[point.X][point.Y] = 1
	}

	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]
		ansX = point.X
		ansY = point.Y

		for i := 0; i < 4; i++ {
			nx := point.X + dir[i][0]
			ny := point.Y + dir[i][1]

			if nx >= 1 && nx <= n && ny >= 1 && ny <= m && vis[nx][ny] == 0 {
				vis[nx][ny] = 1
				queue = append(queue, Point{nx, ny})
			}
		}
	}

	fmt.Fprintln(gout, ansX, ansY)
}
