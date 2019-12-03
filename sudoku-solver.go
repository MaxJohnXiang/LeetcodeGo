// package problem0037

func solveSudoku(board [][]byte) {
	solve(board)
}

/* k 是把 board 转换成一维数组后，元素的索引值 */
func solve(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				for c := '1'; c <= '9'; c++ {
					if isValid(board, i, j, byte(c)) {
						board[i][j] = byte(c)
						if solve(board) {
							return true
						} else {
							board[i][j] = '.'
						}
					}
				}
				return false
			}
		}
	}
	return true
}

func isValid(board [][]byte, row int, col int, num byte) bool {
	blockRow := (row / 3) * 3
	blockCol := (col / 3) * 3

	for i := 0; i < 9; i++ {
		if board[i][col] == num || board[row][i] == num || board[blockRow+i/3][blockCol+i%3] == num {
			return false
		}
	}
	return true
}
