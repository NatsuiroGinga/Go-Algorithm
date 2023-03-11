package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n, k  int
	walls [2][]int
	line  string
)

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Scanln(&n, &k)
	for i := 0; i <= 1; i++ {
		walls[i] = make([]int, n+1)
		line, _ = in.ReadString('\n')

		for j := 1; j <= n; j++ {
			if line[j-1] != 'X' {
				walls[i][j] = 1
			}
		}
	}

	if dfs(1, 0, 1) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func dfs(water, where, height int) bool {
	if height > n {
		return true
	}
	if walls[where][height] == 0 || water > height {
		return false
	}

	walls[where][height] = 0
	return dfs(water+1, where^1, height+k) || dfs(water+1, where, height+1) || dfs(water+1, where, height-1)
}
