package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	fmt.Println(reconstructQueue(nums))
}

func reconstructQueue(people [][]int) [][]int {
	// 先将身高从大到小排序，确定最大个子的相对位置
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1] // 当身高相同时，将K按照从小到大排序
		}
		return people[i][0] > people[j][0] // 身高按照由大到小的顺序来排
	})

	// 再按照K进行插入排序，优先插入K小的
	for i, p := range people {
		copy(people[p[1]+1:i+1], people[p[1]:i+1]) // 空出一个位置
		people[p[1]] = p
	}

	return people
}
