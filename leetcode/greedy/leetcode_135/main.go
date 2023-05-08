package leetcode_135

func candy(ratings []int) int {
	c := make([]int, len(ratings))
	c[0] = 1

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			c[i] = c[i-1] + 1
			continue
		}
		c[i] = 1
	}

	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i] < ratings[i-1] {
			c[i-1] = max(c[i-1], c[i]+1)
		}
	}

	return sum(c)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sum(data []int) int {
	ans := 0
	for _, v := range data {
		ans += v
	}
	return ans
}
