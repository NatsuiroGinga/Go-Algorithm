package day5

import "fmt"

func isValidSudoku(board [][]byte) bool {
	mp := make(map[string]byte, len(board))

	for row, bytes := range board {
		for col, b := range bytes {
			if b != '.' {
				key := fmt.Sprintf("%c in row %d", b, row)
				if _, ok := mp[key]; ok {
					return false
				} else {
					mp[key] = b
				}

				key = fmt.Sprintf("%c in col %d", b, col)
				if _, ok := mp[key]; ok {
					return false
				} else {
					mp[key] = b
				}

				key = fmt.Sprintf("%c in block (%d, %d)", b, row/3, col/3)
				if _, ok := mp[key]; ok {
					return false
				} else {
					mp[key] = b
				}
			}
		}
	}

	return true
}
