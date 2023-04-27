package main

func main() {

}

func merge(intervals [][]int) [][]int {
	var ans [][]int

	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i+1][0] >= intervals[i][1] { // 后一个的开始覆盖了前一个结束
			ans = append(ans, []int{intervals[i][0], intervals[i+1][1]})
		}
	}

	return ans
}
