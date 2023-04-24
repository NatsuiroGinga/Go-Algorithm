package main

import (
	"bufio"
	"fmt"
	"os"
)

var dir = [4][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

var vis [][]bool

func main() {
	gin := bufio.NewReader(os.Stdin)
	gout := bufio.NewWriter(os.Stdout)
	defer gout.Flush()

	var n, m, x, y int
	fmt.Fscan(gin, &n, &m)
	room := make([][]bool, n)
	t := make([][]int, n)
	vis = make([][]bool, n)

	for i := range room {
		room[i] = make([]bool, m)
		t[i] = make([]int, m)
		vis[i] = make([]bool, m)

		for j := range room[i] {
			var mask int
			fmt.Fscan(gin, &mask)
			room[i][j] = mask == 1
			t[i][j] = -1
		}
	}

	fmt.Fscan(gin, &x, &y)
	t[x][y] = 0
	room[x][y] = false

	bfs(room, t, x, y, n, m)

	for i := range t {
		for j := range t[i] {
			fmt.Fprintf(gout, "%d ", t[i][j])
		}
		fmt.Fprintln(gout)
	}
}

func bfs(room [][]bool, t [][]int, x, y, n, m int) {
	var q []Man
	q = append(q, Man{x, y})

	for len(q) > 0 {
		man := q[0]
		q = q[1:]
		vis[man.x][man.y] = true

		for i := range dir {
			nx := man.x + dir[i][0]
			ny := man.y + dir[i][1]

			if nx >= 0 && nx < n && ny >= 0 && ny < m && !vis[nx][ny] {
				// 戴口罩
				if room[nx][ny] {
					// 检查四周
					count := 0
					for k := range dir {
						tx := nx + dir[k][0]
						ty := nx + dir[k][1]
						if tx >= 0 && tx < n && ty >= 0 && ty < m && t[tx][ty] >= 0 {
							count++
							if count == 2 {
								break
							}
						}
					}

					if count < 2 {
						room[nx][ny] = false
					} else {
						t[nx][ny] = t[man.x][man.y] + 1
					}

				} else { // 不戴口罩
					t[nx][ny] = t[man.x][man.y] + 1
					q = append(q, Man{nx, ny})
				}
			}
		}
	}
}

type Man struct {
	x, y int
}
