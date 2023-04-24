package leetcode

func dailyTemperatures(temperatures []int) []int {
	var st []int
	ans := make([]int, len(temperatures))

	for i, t := range temperatures {
		for t > st[len(st)-1] {
			sz := len(st)
			ans[st[sz-1]] = st[sz-1] - i
			st = st[:sz-1]
		}
		st = append(st, i)
	}

	return ans
}
