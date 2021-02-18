package _079_Word_Search

type pair struct{ x, y int }

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

func exist(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])
	vis := make([][]bool, h)
	for i := range vis {
		vis[i] = make([]bool, w)
	}
	var check func(i, j, k int) bool
	check = func(i, j, k int) bool {
		if board[i][j] != word[k] { // 剪枝：当前字符不匹配
			return false
		}
		if k == len(word)-1 { // 单词存在于网格中
			return true
		}
		vis[i][j] = true
		defer func() { vis[i][j] = false }() // 回溯时还原已访问的单元格
		for _, dir := range directions {
			if newI, newJ := i+dir.x, j+dir.y; 0 <= newI && newI < h && 0 <= newJ && newJ < w && !vis[newI][newJ] {
				if check(newI, newJ, k+1) {
					return true
				}
			}
		}
		return false
	}
	for i, row := range board {
		for j := range row {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}

// timeout
func exist1(board [][]byte, word string) bool {
	var ans bool
	for i, line := range board {
		for j, b := range line {
			if b == word[0] {
				visit := make([][]bool, len(board))
				for k := 0; k < len(visit); k++ {
					visit[k] = make([]bool, len(line))
				}
				search(board, word, &visit, i, j, 0, &ans)
				if ans {
					return true
				}
			}
		}
	}
	return false
}

func search(board [][]byte, word string, visit *[][]bool, i, j, index int, finish *bool) {
	if word[index] != board[i][j] {
		return
	}
	(*visit)[i][j] = true
	index++
	if index == len(word) {
		*finish = true
		return
	}
	if i+1 < len(board) && !(*visit)[i+1][j] {
		search(board, word, visit, i+1, j, index, finish)
	}
	if i-1 >= 0 && !(*visit)[i-1][j] {
		search(board, word, visit, i-1, j, index, finish)
	}
	if j+1 < len(board[0]) && !(*visit)[i][j+1] {
		search(board, word, visit, i, j+1, index, finish)
	}
	if j-1 >= 0 && !(*visit)[i][j-1] {
		search(board, word, visit, i, j-1, index, finish)
	}
	(*visit)[i][j] = false
}
