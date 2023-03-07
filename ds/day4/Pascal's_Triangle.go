package day4

func generate(numRows int) [][]int {
	tri := make([][]int, numRows)
	tri[0] = []int{1}

	for i := range tri[1:] {
		i++
		tri[i] = make([]int, i+1)
		tri[i][0], tri[i][i] = 1, 1
		for j := 1; j < i; j++ {
			tri[i][j] = tri[i-1][j-1] + tri[i-1][j]
		}
	}

	return tri
}
