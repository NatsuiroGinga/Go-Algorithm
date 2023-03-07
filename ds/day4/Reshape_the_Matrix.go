package day4

func matrixReshape(mat [][]int, r int, c int) [][]int {
	// 1 <= m, n <= 100
	rows := len(mat)
	cols := len(mat[0])
	if rows*cols != r*c {
		return mat
	}
	ans := make([][]int, r)
	for i := 0; i < r; i++ {
		ans[i] = make([]int, c)
	}

	k := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			ans[i][j] = mat[k/cols][k%cols]
			k++
		}
	}

	return ans
}
