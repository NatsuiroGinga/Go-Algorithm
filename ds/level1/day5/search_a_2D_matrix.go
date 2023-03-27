package day5

func searchMatrix(matrix [][]int, target int) bool {
	var (
		n    = len(matrix)
		m    = len(matrix[0])
		low  = 0
		high = n*m - 1
	)

	for low <= high {
		mid := (low + high) / 2
		row := mid / m
		col := mid % m

		if matrix[row][col] > target {
			high = mid - 1
		} else if matrix[row][col] < target {
			low = mid + 1
		} else {
			return true
		}
	}

	return false
}
