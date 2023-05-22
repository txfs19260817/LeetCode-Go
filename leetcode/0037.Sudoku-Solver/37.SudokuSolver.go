package leetcode

type position struct {
	x, y int
}

func solveSudoku(board [][]byte) {
	var pos []position
	var finish bool
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				pos = append(pos, position{i, j})
			}
		}
	}
	put(&board, pos, 0, &finish)
}

func put(board *[][]byte, pos []position, index int, finish *bool) {
	if *finish || index == len(pos) {
		*finish = true
		return
	}
	for i := 1; i < 10; i++ {
		if check(board, pos[index], byte(i)+'0') && !*finish {
			(*board)[pos[index].x][pos[index].y] = byte(i) + '0'
			put(board, pos, index+1, finish)
			if *finish {
				return
			}
			(*board)[pos[index].x][pos[index].y] = '.'
		}
	}
}

func check(board *[][]byte, p position, v byte) bool {
	for i := 0; i < len(*board); i++ {
		if (*board)[p.x][i] != '.' && (*board)[p.x][i] == v || (*board)[i][p.y] != '.' && (*board)[i][p.y] == v {
			return false
		}
	}
	posx, posy := p.x-p.x%3, p.y-p.y%3
	for i := posx; i < posx+3; i++ {
		for j := posy; j < posy+3; j++ {
			if (*board)[i][j] != '.' && (*board)[i][j] == v {
				return false
			}
		}
	}
	return true
}
