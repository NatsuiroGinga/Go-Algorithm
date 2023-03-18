package day9

import "sort"

/*
Intuition: BFS And DP
Approach:
**First**, use a queue `queue` to store the rotten oranges, use a 2D array `ans` to store the minutes of each minute fresh oranges, which adjacent to rotten oranges, are rotten and use a variable `count` to count the number of fresh oranges.
**Then**, we can use BFS to check the 4 directions of each rotten orange, if there is a fresh orange, we can mark it as rotten,  update the minutes of this fresh orange, `count`--,and push it to the `queue`.
**Finally**, we can check if there are still fresh oranges(count == 0), if so, return -1, otherwise, return the maximum minutes.
*/

type Cell struct {
	x, y int
}

var dirs = [][]int{
	{0, 1},  // right
	{0, -1}, // left
	{1, 0},  // down
	{-1, 0}, // up
}

func orangesRotting(grid [][]int) int {
	var (
		queue   = make([]Cell, 0) // queue
		rows    = len(grid)
		ans     = make([][]int, rows) // answer
		cols    = len(grid[0])
		count   = 0 // count of fresh oranges
		minutes = 0 // minutes
	)

	sort.Reverse()

	for i, line := range grid {
		ans[i] = make([]int, cols)

		for j, v := range line {
			if v == 2 { // rotten
				// push to queue
				queue = append(queue, Cell{i, j})
			} else if v == 1 { // fresh
				count++ // count++
			}
		}
	}

	for len(queue) > 0 {
		// pop
		cell := queue[0]
		queue = queue[1:]

		for _, pair := range dirs { // check 4 directions
			nx := pair[0] + cell.x
			ny := pair[1] + cell.y

			// if not fresh and not out of range
			if nx >= 0 && nx < rows && ny >= 0 && ny < cols && grid[nx][ny] == 1 {
				count--                               // means one fresh orange is rotten
				grid[nx][ny] = 2                      // mark as rotten
				ans[nx][ny] = ans[cell.x][cell.y] + 1 // update answer and push to queue

				// update minutes
				if ans[nx][ny] > minutes {
					minutes = ans[nx][ny]
				}

				// push to queue
				queue = append(queue, Cell{nx, ny})
			}
		}

	}

	// if there are still fresh oranges
	if count != 0 {
		return -1
	}

	return minutes
}
