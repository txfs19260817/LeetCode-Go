package leetcode

func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[i] {
			if search(board, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func search(board [][]byte, word string, index, i, j int) bool {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || word[index] != board[i][j] {
		return false
	}
	if index == len(word)-1 {
		return true
	}

	tmp := board[i][j]
	board[i][j] = 0
	found := search(board, word, index+1, i+1, j) ||
		search(board, word, index+1, i, j+1) ||
		search(board, word, index+1, i-1, j) ||
		search(board, word, index+1, i, j-1)
	board[i][j] = tmp
	return found
}
