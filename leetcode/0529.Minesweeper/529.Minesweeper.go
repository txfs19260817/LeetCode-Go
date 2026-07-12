package leetcode

func updateBoard(board [][]byte, click []int) [][]byte {
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}
	if board[click[0]][click[1]] != 'E' {
		return board
	}
	m, n := len(board), len(board[0])
	dirs := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != 'E' {
			return
		}

		var mines int
		for _, dir := range dirs {
			if ni, nj := i+dir[0], j+dir[1]; ni >= 0 && ni < m && nj >= 0 && nj < n {
				if board[ni][nj] == 'M' {
					mines++
				}
			}
		}

		if mines > 0 {
			board[i][j] = byte('0' + mines)
		} else {
			board[i][j] = 'B'
			for _, dir := range dirs {
				dfs(i+dir[0], j+dir[1])
			}
		}
	}
	dfs(click[0], click[1])
	return board
}

func updateBoard2(board [][]byte, click []int) [][]byte {
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}
	if board[click[0]][click[1]] != 'E' {
		return board
	}
	m, n := len(board), len(board[0])
	dirs := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	q := [][2]int{{click[0], click[1]}}
	for len(q) > 0 {
		i, j := q[0][0], q[0][1]
		q = q[1:]

		if board[i][j] != 'E' {
			continue
		}

		var mines int
		for _, dir := range dirs {
			if ni, nj := i+dir[0], j+dir[1]; ni >= 0 && ni < m && nj >= 0 && nj < n {
				if board[ni][nj] == 'M' {
					mines++
				}
			}
		}

		if mines > 0 {
			board[i][j] = byte('0' + mines)
		} else {
			board[i][j] = 'B'
			for _, dir := range dirs {
				if ni, nj := i+dir[0], j+dir[1]; ni >= 0 && ni < m && nj >= 0 && nj < n {
					q = append(q, [2]int{i + dir[0], j + dir[1]})
				}
			}
		}
	}
	return board
}
