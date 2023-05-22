package leetcode

func isValidSudoku(board [][]byte) bool {
	rows, cols, boxes := [9][9]int{}, [9][9]int{}, [9][9]int{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				continue
			}
			boxIdx, v := i/3*3+j/3, board[i][j]-'1'
			rows[i][v]++
			cols[j][v]++
			boxes[boxIdx][v]++
			if rows[i][v] > 1 || cols[j][v] > 1 || boxes[boxIdx][v] > 1 {
				return false
			}
		}
	}
	return true
}
