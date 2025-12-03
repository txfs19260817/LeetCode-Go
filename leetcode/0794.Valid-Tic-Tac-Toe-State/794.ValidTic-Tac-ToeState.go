package leetcode

import "strings"

func validTicTacToe(board []string) bool {
	win := func(p byte) bool {
		for i := range 3 {
			if board[i][0] == p && board[i][1] == p && board[i][2] == p || board[0][i] == p && board[1][i] == p && board[2][i] == p {
				return true
			}
		}
		return board[0][0] == p && board[1][1] == p && board[2][2] == p || board[0][2] == p && board[1][1] == p && board[2][0] == p
	}
	var no, nx int
	for _, r := range board {
		no += strings.Count(r, "O")
		nx += strings.Count(r, "X")
	}
	return (no == nx || no+1 == nx) && (no == nx || !win('O')) && (no+1 == nx || !win('X'))
}
