package main

import "container/list"

func pacificAtlantic(heights [][]int) [][]int {
	var ans [][]int
	pQueue := list.New()
	aQueue := list.New()
	n, m := len(heights), len(heights[0])
	pacific := make([][]bool, n)
	atlantic := make([][]bool, n)
	for i := 0; i < n; i++ {
		pacific[i] = make([]bool, m)
		atlantic[i] = make([]bool, m)
	}

	// pacific
	for i := 0; i < n; i++ {
		pQueue.PushBack(Cell{i, 0})
		pacific[i][0] = true
		aQueue.PushBack(Cell{i, m - 1})
		atlantic[i][m-1] = true
	}

	// atlantic
	for j := 0; j < m; j++ {
		pQueue.PushBack(Cell{0, j})
		pacific[0][j] = true
		aQueue.PushBack(Cell{n - 1, j})
		atlantic[n-1][j] = true
	}

	bfs(heights, pacific, pQueue)
	bfs(heights, atlantic, aQueue)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if atlantic[i][j] && pacific[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}

	return ans
}

func bfs(heights [][]int, vis [][]bool, queue *list.List) {
	n, m := len(heights), len(heights[0])
	for queue.Len() > 0 {
		cell := queue.Front()
		queue.Remove(cell)

		for _, pair := range dir {
			x, y := cell.Value.(Cell).x, cell.Value.(Cell).y
			nx := pair[0] + x
			ny := pair[1] + y

			if nx >= 0 && nx < n && ny >= 0 && ny < m && !vis[nx][ny] && heights[nx][ny] >= heights[x][y] {
				queue.PushBack(Cell{nx, ny})
				vis[nx][ny] = true
			}
		}
	}
}

type Cell struct {
	x, y int
}
