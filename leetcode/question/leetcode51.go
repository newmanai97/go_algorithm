package question

import "strings"

var (
	res13      [][]string
	chessboard [][]string
)

func solveNQueens(n int) [][]string {
	res13, chessboard = make([][]string, 0), make([][]string, n)
	for i := 0; i < n; i++ {
		chessboard[i] = make([]string, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			chessboard[i][j] = "."
		}
	}
	dfs13(0, n)
	return res13
}

func dfs13(row, n int) {
	if row == n {
		tmp := make([]string, n)
		for i, rowstr := range chessboard {
			tmp[i] = strings.Join(rowstr, "")
		}
		res13 = append(res13, tmp)
		return
	}
	for i := 0; i < n; i++ {
		if isvalid(row, i, n) {
			chessboard[row][i] = "Q"
			dfs13(row+1, n)
			chessboard[row][i] = "."
		}

	}
}

func isvalid(row, col, n int) bool {
	for i := 0; i < row; i++ {
		if chessboard[i][col] == "Q" {
			return false
		}
	}

	for i := 0; i < col; i++ {
		if chessboard[row][i] == "Q" {
			return false
		}
	}

	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}

	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}
	return true
}
