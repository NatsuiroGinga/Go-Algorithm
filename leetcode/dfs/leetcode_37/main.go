package main

func solveSudoku(board [][]byte) {
	backtrack(board)
}

func backtrack(board [][]byte) (canFill bool) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.' {
				for k := byte('1'); k <= '9'; k++ {
					if isValid(i, j, k, board) {
						board[i][j] = k

						if backtrack(board) {
							return true
						}

						board[i][j] = '.'
					}
				}
				return false
			}
		}
	}

	return true
}

func isValid(i, j int, k byte, board [][]byte) bool {
	// 行, 列搜索
	for delta := 0; delta < len(board); delta++ {
		if board[delta][j] == k || board[i][delta] == k {
			return false
		}
	}
	// 九宫格搜索, i/3, j/3
	for r, rcnt := i/3*3, 0; rcnt < 3; r, rcnt = r+1, rcnt+1 {
		for c, ccnt := j/3*3, 0; ccnt < 3; c, ccnt = c+1, ccnt+1 {
			if board[r][c] == k {
				return false
			}
		}
	}

	return true
}

func main() {

}
