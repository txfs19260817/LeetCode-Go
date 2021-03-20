package _130_Surrounded_Regions

func solve(board [][]byte) {
	for i := 0; i < len(board); i++ {
		dfs(board, i, 0)
		dfs(board, i, len(board[0])-1)
	}
	for i := 1; i < len(board[0])-1; i++ {
		dfs(board, 0, i)
		dfs(board, len(board)-1, i)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, x, y int) {
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || board[x][y] != 'O' {
		return
	}
	board[x][y] = 'A'
	dfs(board, x+1, y)
	dfs(board, x-1, y)
	dfs(board, x, y+1)
	dfs(board, x, y-1)
}
