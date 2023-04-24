package main

func main() {
}

func spiralOrder(matrix [][]int) []int {
	var ans []int
	var (
		rowBegin = 0
		rowEnd   = len(matrix) - 1
		colBegin = 0
		colEnd   = len(matrix[0]) - 1
	)

	for colBegin <= colEnd && rowBegin <= rowEnd {
		for i := colBegin; i <= colEnd; i++ {
			ans = append(ans, matrix[rowBegin][i])
		}
		rowBegin++

		for i := rowBegin; i <= rowEnd; i++ {
			ans = append(ans, matrix[i][colEnd])
		}
		colEnd--

		if rowBegin <= rowEnd {
			for i := colEnd; i >= colBegin; i-- {
				ans = append(ans, matrix[rowEnd][i])
			}
		}
		rowEnd--

		if colBegin <= colEnd {
			for i := rowEnd; i >= rowBegin; i-- {
				ans = append(ans, matrix[i][colBegin])
			}
		}
		colBegin++
	}

	return ans
}
