package main

import (
	"sort"
)

func main() {
	p := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	findMinArrowShots(p)
}

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		if points[i][1] == points[j][1] {
			return points[i][0] < points[j][0]
		}
		return points[i][1] < points[j][1]
	})

	bound := points[0][1]
	count := 1
	for i := 1; i < len(points); i++ {
		if points[i][0] < bound {
			continue
		}
		bound = points[i][1]
		count++
	}

	return count
}
