package _051_N_Queens

type position struct {
	x, y int
}

func solveNQueens(n int) [][]string {
	var ans [][]string
	board := generateBoard(n)
	put(&ans, &board, 0, n)
	return ans
}

func put(ans *[][]string, board *[]string, row, n int) {
	if row == n {
		temp := make([]string, n)
		copy(temp, *board)
		*ans = append(*ans, temp)
		return
	}
	for col := 0; col < n; col++ {
		if check(board, position{row, col}, n) {
			bytes := []rune((*board)[row])
			bytes[col] = 'Q'
			(*board)[row] = string(bytes)

			put(ans, board, row+1, n)

			bytes = []rune((*board)[row])
			bytes[col] = '.'
			(*board)[row] = string(bytes)
		}
	}
}

func check(board *[]string, p position, n int) bool {
	for i := 0; i < n; i++ {
		if (*board)[p.x][i] == 'Q' || (*board)[i][p.y] == 'Q' {
			return false
		}
		if p.x-i >= 0 && p.y-i >= 0 && (*board)[p.x-i][p.y-i] == 'Q' {
			return false
		}
		if p.x+i < n && p.y+i < n && (*board)[p.x+i][p.y+i] == 'Q' {
			return false
		}
		if p.x+i < n && p.y-i >= 0 && (*board)[p.x+i][p.y-i] == 'Q' {
			return false
		}
		if p.x-i >= 0 && p.y+i < n && (*board)[p.x-i][p.y+i] == 'Q' {
			return false
		}
	}
	return true
}

func generateBoard(n int) []string {
	board, res := make([]string, 0, n), ""
	for i := 0; i < n; i++ {
		res += "."
	}
	for i := 0; i < n; i++ {
		board = append(board, res)
	}
	return board
}
