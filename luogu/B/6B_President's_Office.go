package B

import (
	"fmt"
)

var (
	n, m   int
	pColor byte
	room   [][]byte
	set    map[byte]struct{}
)

var dir = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func main() {
	fmt.Scanf("%d %d %c\n", &n, &m, &pColor)
	room = make([][]byte, n)
	set = make(map[byte]struct{})
	for i := range room {
		room[i] = make([]byte, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scanf("%c", &room[i][j])
		}
		fmt.Scanln()
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if room[i][j] == pColor {
				for _, d := range dir {
					var nx = i + d[0]
					var ny = j + d[1]
					if nx >= 0 && nx < n && ny >= 0 && ny < m && room[nx][ny] != '.' && room[nx][ny] != pColor {
						set[room[nx][ny]] = struct{}{}
					}
				}
			}
		}
	}

	fmt.Println(len(set))
}
