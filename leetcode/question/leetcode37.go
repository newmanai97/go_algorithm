package question

func solveSudoku(board [][]byte) {
	dfs14(board)
}

func dfs14(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				continue
			}
			for k := '1'; k <= '9'; k++ {
				if isvalid14(i, j, byte(k), board) {
					board[i][j] = byte(k)
					if dfs14(board) {
						return true
					}
					board[i][j] = '.'
				}
			}
			return false
		}
	}
	return true
}

func isvalid14(row, col int, k byte, board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if board[i][col] == k {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		if board[row][i] == k {
			return false
		}
	}
	startrow := (row / 3) * 3
	startcol := (col / 3) * 3
	for i := startrow; i < startrow+3; i++ {
		for j := startcol; j < startcol+3; j++ {
			if board[i][j] == k {
				return false
			}
		}
	}
	return true
}
