package databricks

// DropFigure drops the connected figure ('F') downward as a whole until it
// would collide with an obstacle ('#') or the board bottom, then returns board.
func DropFigure(board [][]byte) [][]byte {
	if len(board) == 0 || len(board[0]) == 0 {
		return board
	}

	rows, cols := len(board), len(board[0])
	minMove := rows
	hasFigure := false
	positions := make([][2]int, 0)

	for c := 0; c < cols; c++ {
		lastF := -1
		for r := 0; r < rows; r++ {
			if board[r][c] == 'F' {
				hasFigure = true
				lastF = r
				positions = append(positions, [2]int{r, c})
			} else if board[r][c] == '#' && lastF != -1 {
				d := r - lastF - 1
				if d < minMove {
					minMove = d
				}
			}
		}
		if lastF != -1 {
			d := rows - lastF - 1
			if d < minMove {
				minMove = d
			}
		}
	}

	if !hasFigure || minMove <= 0 {
		return board
	}

	for _, p := range positions {
		board[p[0]][p[1]] = '.'
	}
	for _, p := range positions {
		board[p[0]+minMove][p[1]] = 'F'
	}

	return board
}
